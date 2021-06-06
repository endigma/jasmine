package main

import (
	"os"

	"gitcat.ca/endigma/jasmine/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	//
	viper.AddConfigPath("$HOME/.config/jasmine")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	// Envars
	viper.SetEnvPrefix("jasmine")

	// Defaults
	viper.SetDefault("initsystem", "runit")

	// Runit defaults
	viper.SetDefault("runit.svdir", "/etc/sv")
	viper.SetDefault("runit.runsvdir", "/var/service")
	viper.SetDefault("runit.timeout", 5)

	viper.AutomaticEnv()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Str("file", viper.ConfigFileUsed()).Msg("Config loaded")
	}

	cmd.Execute()
}
