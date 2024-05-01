package exporter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fr-bez-aosc/znuny-exporter/internal/config"
	"github.com/fr-bez-aosc/znuny-exporter/internal/middlewares"
	"github.com/fr-bez-aosc/znuny-exporter/internal/znuny"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	cfg config.Config

	Data znuny.HealthCheck
)

// register create the exporter specificities
func register() {
	collector := NewHealthCheckCollector()
	prometheus.MustRegister(collector)
}

// Handle set the metrics exposition
func Handle() error {
	http.Handle(cfg.Exporter.Path, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(`
            <html>
            <head><title>Healthcheck status Exporter</title></head>
            <body>
            <h1>Prometheus Znuny's Healthcheck Exporter</h1>
            <p><a href='` + cfg.Exporter.Path + `'>Metrics</a></p>
            </body>
            </html>
        `))
	})
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Exporter.Address, cfg.Exporter.Port), nil)
}

// Serve launch the prometheus exporter
func Serve() {
	// Load the configuration parameters
	if err := cfg.LoadConfig(); err != nil {
		middlewares.TemporaryLogger("Failed to load the configuration file", err)
	}

	// Initialize the logger
	logger := middlewares.NewLogger(cfg.Logs.Path)

	// Run datas recovery
	go func() {
		for {
			// Get healthcheck status
			if err := Data.Get(cfg); err != nil {
				logger.Error(
					"Something went wrong during data recovery",
					zap.Error(err),
				)
			}

			// Wait the next datas recovery
			time.Sleep(time.Duration(cfg.Znuny.Interval) * time.Second)
		}
	}()

	// Register the exporter specificities
	register()

	// Send a log before starting the server
	logger.Info(
		"The exporter is starting",
		zap.String("listenning", fmt.Sprintf("http://%s:%d", cfg.Exporter.Address, cfg.Exporter.Port)),
	)

	// Launch the http server to expose metrics
	if err := Handle(); err != nil {
		logger.Error(
			"Something went wrong with the web server",
			zap.Error(err),
		)
	}
}
