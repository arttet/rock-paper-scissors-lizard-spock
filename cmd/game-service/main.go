package main

import (
	"flag"
	"log"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/app/game-service/server"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/config"
	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/telemetry"

	"go.uber.org/zap"
)

func main() {
	configYML := flag.String("cfg", "config.yml", "Defines the configuration file option")
	flag.Parse()

	cfg, err := config.NewConfigWithYAML(*configYML)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := cfg.Logger.Build()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(logger)
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Fatal(err)
		}
	}()

	logger.Info("starting service",
		zap.String("name", cfg.Project.Name),
		zap.String("yml", *configYML),
		zap.Bool("debug", cfg.Project.Debug),
		zap.String("environment", cfg.Project.Environment),
		zap.String("commit_hash", cfg.Project.CommitHash),
		zap.String("version", cfg.Project.Version),
	)

	tracing, err := telemetry.NewTracer(cfg.Jaeger)
	if err != nil {
		logger.Fatal("tracing initialization", zap.Error(err))
	}
	defer tracing.Close()

	if err := server.NewServer(logger).Start(cfg); err != nil {
		logger.Error("server initialization", zap.Error(err))
	}
}
