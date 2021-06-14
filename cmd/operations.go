package cmd

import (
	"gitcat.ca/endigma/jasmine/util"
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
				if err := initSystem.Add(name); err != nil {
					log.Fatal().Msg(err.Error())
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
				if err := initSystem.Remove(name); err != nil {
					log.Fatal().Msg(err.Error())
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
				if err := initSystem.Enable(name); err != nil {
					log.Fatal().Msg(err.Error())
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
				if err := initSystem.Disable(name); err != nil {
					log.Fatal().Msg(err.Error())
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
				if err := initSystem.Start(name); err != nil {
					log.Fatal().Msg(err.Error())
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
				if err := initSystem.Stop(name); err != nil {
					log.Fatal().Msg(err.Error())
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
				if err := initSystem.Restart(name); err != nil {
					log.Fatal().Msg(err.Error())
				}
			}
		},
	}
	cmd_operation_once = &cobra.Command{
		Use:   "once [services...]",
		Short: "Once named services",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			for _, name := range args {
				if err := initSystem.Once(name); err != nil {
					log.Fatal().Msg(err.Error())
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
				if err := initSystem.Reload(name); err != nil {
					log.Fatal().Msg(err.Error())
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
