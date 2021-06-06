package main

import (
	"os"
	"os/user"

	"gitcat.ca/endigma/jasmine/cmd"
	"gitcat.ca/endigma/jasmine/util"
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
	viper.BindEnv("suppress_permissions_warning")

	// Defaults
	viper.SetDefault("initsystem", "runit")
	viper.SetDefault("runit.svdir", "/etc/sv")
	viper.SetDefault("runit.runsvdir", "/var/service")
	viper.SetDefault("runit.timeout", 5)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Str("file", viper.ConfigFileUsed()).Msg("Config loaded")
	}

	if !func() bool {
		currentUser, err := user.Current()
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to get current user")
		}
		return currentUser.Username == "root"
	}() {
		if !viper.GetBool("suppress_permissions_warning") {
			util.FancyWarn("Jasmine must be run as root for most commands to work\n\n")
		}
	}

	cmd.Execute()
}
