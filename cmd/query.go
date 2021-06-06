package cmd

import (
	"fmt"
	"os"
	"time"

	"gitcat.ca/endigma/jasmine/inits"
	"github.com/fatih/color"
	"github.com/hako/durafmt"
	"github.com/juju/ansiterm"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	cmd_query_list = &cobra.Command{
		Use:     "list [filter services...]",
		Short:   "List all running services",
		Aliases: []string{"ls", "ll", "show"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			var list []inits.Service

			list, err := initSystem.List(args)
			if err != nil {
				log.Fatal().Msg(err.Error())
			}

			w := ansiterm.NewTabWriter(os.Stdout, 1, 1, 1, ' ', 0)
			color.New(color.Bold).Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", "SERVICE", "STATE", "ENABLED", "PID", "COMMAND", "TIME")
			for _, f := range list {
				fmt.Fprintf(w, "%s\t%s\t%v\t%s\t%.17s\t%s\n",
					f.Name,
					func(f inits.Service) string {
						switch f.State {
						case "up":
							return color.New(color.FgGreen).Sprint(f.State)
						case "down":
							return color.New(color.FgRed).Sprint(f.State)
						}
						return ""
					}(f),
					func(f inits.Service) string {
						switch f.State {
						case "up":
							return color.New(color.FgGreen).Sprint(f.Enabled)
						case "down":
							return color.New(color.FgRed).Sprint(f.Enabled)
						}
						return ""
					}(f),
					func(f inits.Service) string {
						switch f.State {
						case "up":
							return color.New(color.FgHiMagenta).Sprint(f.PID)
						case "down":
							return color.New(color.FgHiBlack).Sprint("---")
						}
						return ""
					}(f),
					func(f inits.Service) string {
						switch f.State {
						case "up":
							return f.Command
						case "down":
							return color.New(color.FgHiBlack).Sprint("---")
						}
						return ""
					}(f),
					func(f inits.Service) string {
						switch true {
						case f.Uptime < 5*time.Minute:
							return color.New(color.FgRed).Sprint(durafmt.Parse(f.Uptime).LimitFirstN(1))
						case f.Uptime < 30*time.Minute:
							return color.New(color.FgYellow).Sprint(durafmt.Parse(f.Uptime).LimitFirstN(1))
						}
						return color.New(color.FgHiBlack).Sprint(durafmt.Parse(f.Uptime).LimitFirstN(1))
					}(f),
				)
			}

			w.Flush()
		},
	}
	cmd_query_list_available = &cobra.Command{
		Use:     "listavailable [filter services...]",
		Short:   "List all available services",
		Aliases: []string{"la"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			var list map[string]bool

			list, err := initSystem.ListAvailable()
			if err != nil {
				log.Fatal().Msg(err.Error())
			}

			for servicefile, enabled := range list {
				if enabled {
					color.New(color.FgGreen).Println(servicefile)
				} else {
					color.New(color.FgHiBlack).Println(servicefile)
				}

			}
		},
	}
)

func init() {
	cmd_root.AddCommand(cmd_query_list, cmd_query_list_available)
}
