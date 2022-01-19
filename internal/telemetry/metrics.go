package telemetry

import (
	"fmt"
	"net/http"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewMetricsServer returns a new metrics server.
func NewMetricsServer(cfg config.Metrics) *http.Server {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	mux := http.DefaultServeMux
	mux.Handle(cfg.Path, promhttp.Handler())

	metricsServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return metricsServer
}
