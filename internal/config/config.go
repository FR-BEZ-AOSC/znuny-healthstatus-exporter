package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config define the global structure of the exporter configuration
type Config struct {
	Exporter exporter `mapstructure:"exporter"`
	Znuny    znuny    `mapstructure:"znuny"`
}

// LoadConfig manage all settings passed to the exporter
func (c *Config) LoadConfig() (*Config, error) {
	// Configure Viper to search for config.yaml in /etc/nergalio/ and current directory
	viper.SetConfigName("znuny-exporter")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")

	if len(viper.GetString("config-path")) > 0 {
		viper.SetConfigFile(viper.GetString("config-path"))
	}

	// Read configurations from file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal configurations
	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}

	// Automatic creation of the environment
	viper.AutomaticEnv()

	return c, nil
}

// Print display the configuration parameters for debug
func (c *Config) Print() {
	fmt.Printf("\nConfiguration options:\n%v\n\n", c)
}