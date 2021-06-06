package cmd

import (
	"gitcat.ca/endigma/jasmine/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	cmd_operation_enable = &cobra.Command{
		Use:     "enable services...",
		Short:   "Enable named services",
		Aliases: []string{"e"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Enable(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_disable = &cobra.Command{
		Use:     "disable services...",
		Short:   "Disable named services",
		Aliases: []string{"d"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Disable(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_up = &cobra.Command{
		Use:   "up services...",
		Short: "Up named services",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Up(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_down = &cobra.Command{
		Use:     "down services...",
		Short:   "Down named services",
		Aliases: []string{"dn"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Down(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_start = &cobra.Command{
		Use:   "start services...",
		Short: "Start named services",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Start(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_stop = &cobra.Command{
		Use:   "stop services...",
		Short: "Stop named services",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Stop(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_restart = &cobra.Command{
		Use:     "restart services...",
		Short:   "Restart named services",
		Aliases: []string{"rs"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Restart(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_once = &cobra.Command{
		Use:     "once services...",
		Short:   "Once named services",
		Aliases: []string{"o"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Once(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_reload = &cobra.Command{
		Use:     "reload services...",
		Short:   "Reload named services",
		Aliases: []string{"o"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			if err := initSystem.Reload(args); err != nil {
				log.Fatal().Msg(err.Error())
			}
		},
	}
	cmd_operation_pass = &cobra.Command{
		Use:     "pass args..",
		Short:   "Pass commands onto your init system's default tool",
		Aliases: []string{"p"},
		Args:    cobra.MinimumNArgs(1),
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
		cmd_operation_enable,
		cmd_operation_disable,
		cmd_operation_up,
		cmd_operation_down,
		cmd_operation_start,
		cmd_operation_stop,
		cmd_operation_restart,
		cmd_operation_once,
		cmd_operation_reload,
		cmd_operation_pass,
	)
}
