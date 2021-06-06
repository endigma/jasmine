package cmd

import (
	"fmt"
	"os"

	"gitcat.ca/endigma/jasmine/inits"
	"gitcat.ca/endigma/jasmine/inits/runit"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initSystem inits.Init

// cmd_root represents the base command when called without any subcommands
var cmd_root = &cobra.Command{
	Use:   os.Args[0],
	Short: "a better user interface for runit",
	Long:  fmt.Sprintf("%[1]s:\n  %[1]s is a cli app that interfaces with runit\n", os.Args[0]),
}

// Execute starts jasmine
func Execute() {
	switch viper.GetString("initsystem") {
	case "runit":
		initSystem = runit.New(viper.GetString("runit.svdir"), viper.GetString("runit.runsvdir"), viper.GetInt("runit.timeout"))
	default:
		log.Fatal().Msg("Unsupported Init System!")
	}

	cmd_root.PersistentFlags().Bool("suppress", false, "Suppress warnings when UID is not 0")
	viper.GetViper().BindPFlag("suppress_permissions_warning", cmd_root.PersistentFlags().Lookup("suppress"))

	cobra.CheckErr(cmd_root.Execute())
}
