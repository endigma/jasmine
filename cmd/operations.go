package cmd

import (
	"fmt"

	"gitcat.ca/endigma/jasmine/util"
	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	cmd_operation_add = &cobra.Command{
		Use:     "add [services...]",
		Short:   "Add named services",
		Aliases: []string{"a"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Adding %s... ", name)
				if err := initSystem.Add(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_remove = &cobra.Command{
		Use:     "remove [services...]",
		Short:   "Remove named services",
		Aliases: []string{"r"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Removing %s... ", name)
				if err := initSystem.Remove(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_enable = &cobra.Command{
		Use:     "enable [services...]",
		Short:   "Enable named services",
		Aliases: []string{"e"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Enabling %s... ", name)
				if err := initSystem.Enable(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_disable = &cobra.Command{
		Use:     "disable [services...]",
		Short:   "Disable named services",
		Aliases: []string{"d"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Disabling %s... ", name)
				if err := initSystem.Disable(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_start = &cobra.Command{
		Use:   "start [services...]",
		Short: "Start named services",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Starting %s... ", name)
				if err := initSystem.Start(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_stop = &cobra.Command{
		Use:   "stop [services...]",
		Short: "Stop named services",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Stopping %s... ", name)
				color.New(color.FgRed).Print("Error!\n")
				if err := initSystem.Stop(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_restart = &cobra.Command{
		Use:   "restart [services...]",
		Short: "Restart named services",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Restarting %s... ", name)
				if err := initSystem.Restart(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_once = &cobra.Command{
		Use:   "once [services...]",
		Short: "Run named services once",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Running %s... ", name)
				if err := initSystem.Once(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_reload = &cobra.Command{
		Use:   "reload [services...]",
		Short: "Reload named services",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				fmt.Printf("Removing %s... ", name)
				if err := initSystem.Reload(name); err != nil {
					color.New(color.FgRed).Printf("Error! %s\n", err.Error())
				} else {
					color.New(color.FgGreen).Print("Done!\n")
				}
			}
		},
	}
	cmd_operation_pass = &cobra.Command{
		Use:   "pass [args...]",
		Short: "Pass commands onto your init system's default tool",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Pass(args...); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
)

func init() {
	cmd_root.AddCommand(
		cmd_operation_add,
		cmd_operation_remove,
		cmd_operation_enable,
		cmd_operation_disable,
		cmd_operation_start,
		cmd_operation_stop,
		cmd_operation_restart,
		cmd_operation_once,
		cmd_operation_reload,
		cmd_operation_pass,
	)
}
