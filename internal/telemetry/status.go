package telemetry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"sync/atomic"

	// nolint:gosec
	_ "net/http/pprof"

	"github.com/arttet/rock-paper-scissors-lizard-spock/internal/config"

	"go.uber.org/zap"
)

// NewStatusServer returns a new status server.
func NewStatusServer(cfg *config.Config, isReady *atomic.Value) *http.Server {
	statusAddr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)

	mux := http.DefaultServeMux

	mux.HandleFunc(cfg.Status.LivenessPath, livenessHandler)
	mux.HandleFunc(cfg.Status.ReadinessPath, readinessHandler(isReady))
	mux.HandleFunc(cfg.Status.VersionPath, versionHandler(cfg))

	mux.HandleFunc(cfg.Status.LoggerPath, cfg.Logger.Level.ServeHTTP)
	mux.HandleFunc(cfg.Status.SwaggerPath, swaggerHandler(cfg))

	statusServer := &http.Server{
		Addr:    statusAddr,
		Handler: mux,
	}

	return statusServer
}

func livenessHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)

			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func versionHandler(cfg *config.Config) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		data := map[string]interface{}{
			"name":        cfg.Project.Name,
			"debug":       cfg.Project.Debug,
			"environment": cfg.Project.Environment,
			"version":     cfg.Project.Version,
			"commitHash":  cfg.Project.CommitHash,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			zap.L().Error("Service information encoding error", zap.Error(err))
		}
	}
}

func swaggerHandler(cfg *config.Config) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := zap.S()

		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			logger.Errorw("swagger JSON not found %v", r.URL.Path)
			http.NotFound(w, r)

			return
		}

		logger.Infow("Serving %s", r.URL.Path)

		filepath := strings.TrimPrefix(r.URL.Path, cfg.Status.SwaggerPath)
		filepath = path.Join(cfg.Status.SwaggerDir, filepath)

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			http.NotFound(w, r)

			return
		}

		http.ServeFile(w, r, filepath)
	}
}
