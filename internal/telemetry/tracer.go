package telemetry

import (
	"fmt"
	"io"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/config"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer returns a new tracer.
func NewTracer(cfg config.Jaeger) (io.Closer, error) {
	tracerAddr := fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)

	tracerCfg := &jaegercfg.Configuration{
		ServiceName: cfg.Service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: tracerAddr,
		},
	}

	logger := zap.L()

	tracer, closer, err := tracerCfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		logger.Fatal("failed to init Jaeger", zap.Error(err))

		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)
	logger.Info("tracer started")

	return closer, nil
}
