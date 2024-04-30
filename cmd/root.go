package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fr-bez-aosc/znuny-exporter/internal/exporter"
)

var rootCmd = &cobra.Command{
	Use:   "znuny-exporter",
	Short: "Prometheus Exporter for the application Znuny",
	Long:  "Prometheus Exporter for the application Znuny wich collecte healthcheck status of an defined instance of Znuny then expose its for a Pormetheus scrapping",
	Run: func(cmd *cobra.Command, args []string) {
		exporter.Serve()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config-path", "c", "", "Config file path")
	viper.BindPFlag("config-path", rootCmd.PersistentFlags().Lookup("config-path"))
	viper.BindEnv("config-path", "ZE_CONFIG_PATH")

	/*
		Exporter configurations
	*/

	rootCmd.Flags().String("exporter-address", "", "Exporter address to listen to")
	viper.BindPFlag("exporter.address", rootCmd.Flags().Lookup("exporter-address"))
	viper.BindEnv("exporter.address", "ZE_EXPORTER_ADDRESS")

	rootCmd.Flags().String("exporter-port", "", "Exporter port to listen to")
	viper.BindPFlag("exporter.port", rootCmd.Flags().Lookup("exporter-port"))
	viper.BindEnv("exporter.port", "ZE_EXPORTER_PORT")

	rootCmd.Flags().String("exporter-path", "", "Exporter path to expose metrics")
	viper.BindPFlag("exporter.path", rootCmd.Flags().Lookup("exporter-path"))
	viper.BindEnv("exporter.path", "ZE_EXPORTER_PATH")

	rootCmd.Flags().String("exporter-interval", "", "Exporter interval in seconds to sent healthcheck requests")
	viper.BindPFlag("exporter.interval", rootCmd.Flags().Lookup("exporter-interval"))
	viper.BindEnv("exporter.interval", "ZE_EXPORTER_INTERVAL")

	/*
		Znuny configurations
	*/

	rootCmd.Flags().String("znuny-domain", "", "Domain of the remote Znuny instance")
	viper.BindPFlag("znuny.domain", rootCmd.Flags().Lookup("znuny-domain"))
	viper.BindEnv("znuny.domain", "ZE_ZNUNY_DOMAIN")

	rootCmd.Flags().String("znuny-token", "", "Token used to request on the remote Znuny instance")
	viper.BindPFlag("znuny.token", rootCmd.Flags().Lookup("znuny-token"))
	viper.BindEnv("znuny.token", "ZE_ZNUNY_TOKEN")

	rootCmd.Flags().String("znuny-tls", "", "Whether or not the https protocol is used")
	viper.BindPFlag("znuny.tls", rootCmd.Flags().Lookup("znuny-tls"))
	viper.BindEnv("znuny.tls", "ZE_ZNUNY_TLS")
}
