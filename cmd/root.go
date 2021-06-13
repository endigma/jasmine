package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"gitcat.ca/endigma/jasmine/inits"
	"gitcat.ca/endigma/jasmine/inits/runit"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initSystem inits.Init

var initSystems = map[string]func() inits.Init{
	"runit": runit.New,
}

// cmd_root represents the base command when called without any subcommands
var cmd_root = &cobra.Command{
	Use:   os.Args[0],
	Short: "a better user interface for init systems and service supervisors",
	Long:  fmt.Sprintf("%[1]s:\n  %[1]s Jasmine is a frontend for init systems like runit, openrc, s6 and systemd.\n", os.Args[0]),
}

// Execute starts jasmine
func Execute() {
	// Config Paths
	viper.AddConfigPath(filepath.Join("/home", os.Getenv("SUDO_USER"), ".config/jasmine"))
	viper.AddConfigPath("/etc/jasmine")
	viper.AddConfigPath(".")
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

	cmd_root.PersistentFlags().Bool("debug", false, "Show debug information")
	cmd_root.PersistentFlags().BoolP("suppress", "s", false, "Suppress warnings when UID is not 0")
	viper.GetViper().BindPFlag("show_debug_info", cmd_root.PersistentFlags().Lookup("debug"))
	viper.GetViper().BindPFlag("suppress_permissions_warning", cmd_root.PersistentFlags().Lookup("suppress"))

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if viper.GetBool("show_debug_info") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Debug().Str("file", viper.ConfigFileUsed()).Msg("Config loaded")
	}

	if initSystemNew, ok := initSystems[viper.GetString("initsystem")]; ok {
		initSystem = initSystemNew()
	} else {
		log.Fatal().Msg("Unsupported Init System!")
	}

	cobra.CheckErr(cmd_root.Execute())
}
