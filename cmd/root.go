package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitcat.ca/endigma/jasmine/inits"
	"gitcat.ca/endigma/jasmine/inits/runit"
	"github.com/fatih/color"
	"github.com/juju/ansiterm"
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
	Use:     fmt.Sprintf("%s [command] [args...] [flags]", os.Args[0]),
	Version: "v0.1.6",
	Short:   "is a better user interface for init systems and service supervisors",
	Long:    fmt.Sprintf("%[1]s:\n  %[1]s Jasmine is a frontend for init systems like runit, openrc, s6 and systemd.\n", os.Args[0]),
}

func help(cmd *cobra.Command, args []string) {
	if cmd == cmd_root {
		var buf *bufio.Writer = bufio.NewWriter(os.Stdout)
		fmt.Fprintf(buf,
			"%[1]s:\n  Jasmine is a frontend for init systems like runit, openrc, s6 and systemd.\n\n",
			color.New(color.FgHiMagenta).Sprint("Jasmine"))

		fmt.Fprintf(buf,
			"%s:\n  %s\n\n",
			color.New(color.FgBlue).Sprint("Usage"), cmd.Use)

		fmt.Fprintf(buf,
			"%s:\n",
			color.New(color.FgGreen).Sprint("Available Commands"))

		w := ansiterm.NewTabWriter(buf, 1, 1, 4, ' ', 0)
		for _, cmd := range cmd.Commands() {
			fmt.Fprintf(w, "  %s\t%s\t%s\n", strings.Split(cmd.Use, " ")[0], color.New(color.FgRed, color.Bold).Sprint(strings.Join(cmd.Aliases, ", ")), cmd.Short)
		}

		fmt.Fprint(w, "\n")

		w.Flush()

		fmt.Fprintf(buf,
			"%s:\n",
			color.New(color.FgHiYellow).Sprint("Global Flags"))

		fmt.Fprint(buf, cmd.LocalFlags().FlagUsages())

		fmt.Fprintf(buf, "\nUse \"%s [command] %s\" for more information about a command.\n", os.Args[0], color.New(color.FgHiYellow).Sprint("--help"))

		buf.Flush()
	} else {
		var buf *bufio.Writer = bufio.NewWriter(os.Stdout)
		fmt.Fprintf(buf,
			"%s:\n  %s\n\n",
			color.New(color.Bold).Sprint(cmd.Name()), cmd.Short)

		if cmd.Long != "" {
			fmt.Printf(
				"%s:\n  %s\n\n",
				color.New(color.FgHiBlue).Sprint("Desc"), cmd.Long)
		}

		fmt.Fprintf(buf,
			"%s:\n  %s\n\n",
			color.New(color.FgBlue).Sprint("Usage"), cmd.Use)

		if len(cmd.Aliases) != 0 {
			fmt.Fprintf(buf,
				"%s:\n  %s\n\n",
				color.New(color.FgRed).Sprint("Aliases"), strings.Join(cmd.Aliases, ", "))
		}

		if cmd.LocalFlags().HasFlags() {
			fmt.Fprintf(buf,
				"%s:\n",
				color.New(color.FgYellow).Sprint("Flags"))

			fmt.Fprint(buf, cmd.LocalFlags().FlagUsages(), "\n")
		}

		if cmd.LocalFlags().HasFlags() {
			fmt.Fprintf(buf,
				"%s:\n",
				color.New(color.FgHiYellow).Sprint("Global Flags"))

			fmt.Fprint(buf, cmd.InheritedFlags().FlagUsages(), "\n")
		}

		buf.Flush()
	}

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

	cmd_root.SetHelpFunc(help)
	cmd_root.SetUsageFunc(func(c *cobra.Command) error {
		help(c, []string{c.Name()})
		return nil
	})

	cmd_root.SetVersionTemplate("jasmine {{printf \"%s\" .Version}}\nauthor: endigma <endigma@mailcat.ca>\nlicense: AGPLv3\nsource: https://gitcat.ca/endigma/jasmine\n")

	for _, cmd := range cmd_root.Commands() {
		cmd.Flags().BoolP("help", "h", false, fmt.Sprintf("Help for %s", cmd.Name()))
	}

	cobra.CheckErr(cmd_root.Execute())
}
