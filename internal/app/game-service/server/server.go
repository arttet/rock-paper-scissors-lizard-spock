package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/api"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service/random"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/service/rpsls"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/config"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/server"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/telemetry"

	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	pb "github.com/arttet/rock-paper-scissors-lizard-spock/pkg/game-service/v1"
)

const (
	n    = 101
	seed = 0
)

type Server struct {
	logger *zap.Logger
}

func NewServer(logger *zap.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Start(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gatewayAddr := fmt.Sprintf("%s:%v", cfg.REST.Host, cfg.REST.Port)
	grpcAddr := fmt.Sprintf("%s:%v", cfg.GRPC.Host, cfg.GRPC.Port)
	metricsAddr := fmt.Sprintf("%s:%v", cfg.Metrics.Host, cfg.Metrics.Port)
	statusAdrr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)

	logger := s.logger
	gatewayServer := newGatewayServer(cfg.REST, grpcAddr, gatewayAddr)

	go func() {
		logger.Info("gateway server is running", zap.String("address", gatewayAddr))
		if err := gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed running gateway server", zap.Error(err))
			cancel()
		}
	}()

	metricsServer := telemetry.NewMetricsServer(cfg.Metrics)

	go func() {
		logger.Info("metrics server is running", zap.String("address", metricsAddr))
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed running metrics server", zap.Error(err))
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)
	statusServer := telemetry.NewStatusServer(cfg, isReady)

	go func() {
		logger.Info("status server is running", zap.String("address", statusAdrr))
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed running status server", zap.Error(err))
			// cancel()
		}
	}()

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: cfg.GRPC.MaxConnectionIdle,
			Timeout:           cfg.GRPC.Timeout,
			MaxConnectionAge:  cfg.GRPC.MaxConnectionAge,
			Time:              cfg.GRPC.Timeout,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_opentracing.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
			server.ValidateInterceptor(),
		)),
	)

	rg := random.NewRandomGenerator(n, seed, logger)
	game := rpsls.NewRockPaperScissorsLizardSpockGame(rg, logger)

	pb.RegisterGameServiceServer(
		grpcServer,
		api.NewGameServiceServer(game, logger),
	)

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)

	go func() {
		logger.Info("gRPC server is running", zap.String("address", grpcAddr))
		if err := grpcServer.Serve(listener); err != nil {
			logger.Error("failed running gRPC server", zap.Error(err))
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		logger.Info("the service is ready to accept requests")
	}()

	if cfg.Project.Debug {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.Info("", zap.String("signal.Notify", fmt.Sprintf("%v", v)))
	case done := <-ctx.Done():
		logger.Info("", zap.String("ctx.Done", fmt.Sprintf("%v", done)))
	}

	isReady.Store(false)

	if err := gatewayServer.Shutdown(ctx); err != nil {
		logger.Error("gateway server shut down", zap.Error(err))
	} else {
		logger.Info("gateway server shut down correctly")
	}

	if err := statusServer.Shutdown(ctx); err != nil {
		logger.Error("status server shut down", zap.Error(err))
	} else {
		logger.Info("status server shut down correctly")
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		logger.Error("metrics server shut down", zap.Error(err))
	} else {
		logger.Info("metrics server shut down correctly")
	}

	grpcServer.GracefulStop()
	logger.Info("gRPC server shut down correctly")

	return nil
}
