package exporter

import (
	cfg "github.com/fr-bez-aosc/znuny-exporter/internal/config"
	api "github.com/fr-bez-aosc/znuny-exporter/internal/znuny"
)

var (
	c cfg.Config
)

// Serve launch the prometheus exporter
func Serve() {
	// Load the configuration parameters
	config, err := c.LoadConfig()
	if err != nil {
		panic(err)
	}

	// // Print the configurations for debug
	// config.Print()

	// Get healthcheck status
	data, err := api.SendRequest(config.Znuny.Domain, config.Znuny.Token)
	if err != nil {
		panic(err)
	}

	// Print the response datas for debug
	api.Format(data)

	// // Launch the http server to expose metrics
	// http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(fmt.Sprintf("%s:%d", config.Exporter.Address, config.Exporter.Port), nil)
}
