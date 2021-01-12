package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Verbose bool
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "bigbrother",
		Short: "BigBrother is a very fast web analytics server",
		Long: `A Fast and Flexible web analytics server.
Complete documentation is available at http://...`,
		TraverseChildren: true,
	}
)

func er(msg interface{}) {
	log.Error(msg)
	os.Exit(1)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.bb.yaml)")
	viper.SetDefault("license", "apache")
	viper.SetDefault("ch.url", "http://127.127.127.127:8123/default")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		home := filepath.Dir(ex)

		viper.AddConfigPath(home)
		viper.SetConfigName(".bb")
	}

	viper.SetEnvPrefix("bb")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	} else {
		er(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
		os.Exit(1)
	}
}
