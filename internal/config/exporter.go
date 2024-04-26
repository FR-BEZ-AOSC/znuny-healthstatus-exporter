package config

// exporter define the configuration structure for the exporter
type exporter struct {
	Address string `mapstructure:"address"`
	Port int64 `mapstructure:"port"`
}