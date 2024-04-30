package config

// znuny define the configuration structure for the remote Znuny instance
type znuny struct {
	Domain   string `mapstructure:"domain"`
	Token    string `mapstructure:"token"`
	TLS      bool   `mapstructure:"tls"`
	Interval int    `mapstructure:"interval"`
}
