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
func (c *Config) LoadConfig() error {
	// Configure Viper to search for znuny-exporter.yaml in /etc/
	viper.SetConfigName("znuny-exporter")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")

	// If the flag '--config-path' is set, get its value as the configuration file
	if len(viper.GetString("config-path")) > 0 {
		viper.SetConfigFile(viper.GetString("config-path"))
	}

	// Read configurations from file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Unmarshal configurations
	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	// Automatic creation of the environment
	viper.AutomaticEnv()

	return nil
}

// Print display the configuration parameters for debug
func (c *Config) Print() {
	fmt.Printf("\nConfiguration options:\n%v\n\n", c)
}
