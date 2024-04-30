package exporter

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fr-bez-aosc/znuny-exporter/internal/config"
	"github.com/fr-bez-aosc/znuny-exporter/internal/znuny"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
            <head><title>Volume Exporter Metrics</title></head>
            <body>
            <h1>Prometheus Updates Exporter</h1>
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
		panic(err)
	}

	// Run datas recovery
	go func() {
		for {
			// Get healthcheck status
			if err := Data.Get(cfg); err != nil {
				panic(err)
			}

			// Wait the next datas recovery
			time.Sleep(time.Duration(cfg.Znuny.Interval) * time.Second)
		}
	}()

	// Register the exporter specificities
	register()

	// Launch the http server to expose metrics
	log.Printf("znuny_exporter starting on http://%s:%d%s", cfg.Exporter.Address, cfg.Exporter.Port, cfg.Exporter.Path)
	log.Fatal(Handle())
}
