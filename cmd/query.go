package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gitcat.ca/endigma/jasmine/inits"
	"gitcat.ca/endigma/jasmine/util"
	"github.com/fatih/color"
	"github.com/hako/durafmt"
	"github.com/juju/ansiterm"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var scriptable bool

var (
	cmd_query_list = &cobra.Command{
		Use:     "list [filter services...]",
		Short:   "List all running services",
		Aliases: []string{"ls", "ll"},
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()
			var list []inits.Service

			list, err := initSystem.List(args)
			if err != nil {
				log.Fatal().Msg(err.Error())
			}

			w := ansiterm.NewTabWriter(os.Stdout, 1, 1, 1, ' ', 0)
			color.New(color.Bold).Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", "SERVICE", "STATE", "ENABLED", "PID", "COMMAND", "TIME")
			for _, sv := range list {
				fmt.Fprintf(w, "%s\t%s\t%v\t%s\t%.17s\t%s\n",
					sv.Name,
					func(sv inits.Service) string {
						switch sv.State {
						case "up":
							return color.New(color.FgGreen).Sprint(sv.State)
						case "down":
							return color.New(color.FgRed).Sprint(sv.State)
						}
						return ""
					}(sv),
					func(sv inits.Service) string {
						switch sv.Enabled {
						case true:
							return color.New(color.FgGreen).Sprint(sv.Enabled)
						case false:
							return color.New(color.FgRed).Sprint(sv.Enabled)
						}
						return ""
					}(sv),
					func(sv inits.Service) string {
						switch sv.State {
						case "up":
							return color.New(color.FgHiMagenta).Sprint(sv.PID)
						case "down":
							return color.New(color.FgHiBlack).Sprint("---")
						}
						return ""
					}(sv),
					func(sv inits.Service) string {
						switch sv.State {
						case "up":
							return sv.Command[0]
						case "down":
							return color.New(color.FgHiBlack).Sprint("---")
						}
						return ""
					}(sv),
					func(sv inits.Service) string {
						switch true {
						case sv.Uptime < 5*time.Minute:
							return color.New(color.FgRed).Sprint(durafmt.Parse(sv.Uptime).LimitFirstN(1))
						case sv.Uptime < 30*time.Minute:
							return color.New(color.FgYellow).Sprint(durafmt.Parse(sv.Uptime).LimitFirstN(1))
						}
						return color.New(color.FgHiBlack).Sprint(durafmt.Parse(sv.Uptime).LimitFirstN(1))
					}(sv),
				)
			}

			w.Flush()
		},
	}
	cmd_query_list_available = &cobra.Command{
		Use:     "listavail [filter services...]",
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
	cmd_query_status = &cobra.Command{
		Use:     "status [services...]",
		Short:   "Show detailed information about a service",
		Aliases: []string{"s"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.SudoWarn()

			for _, name := range args {
				sv, err := initSystem.Status(name)
				if err != nil {
					log.Fatal().Msg(err.Error())
				}

				if scriptable {
					fmt.Printf("%s,%s,%s,%v,%d,%s\n",
						sv.Name,
						sv.State,
						sv.Uptime.String(),
						sv.Enabled,
						sv.PID,
						strings.Join(sv.Command, " "),
					)
				} else {
					fmt.Printf("%s:\n  uptime: %s\n  status: %s\n  enabled: %s\n  pid: %s\n  command: %s\n",
						sv.Name,
						func(sv inits.Service) string {
							switch true {
							case sv.Uptime < 5*time.Minute:
								return color.New(color.FgRed).Sprint(durafmt.Parse(sv.Uptime).LimitFirstN(4))
							case sv.Uptime < 30*time.Minute:
								return color.New(color.FgYellow).Sprint(durafmt.Parse(sv.Uptime).LimitFirstN(4))
							}
							return color.New(color.FgHiBlack).Sprint(durafmt.Parse(sv.Uptime).LimitFirstN(4))
						}(sv),
						func(sv inits.Service) string {
							switch sv.State {
							case "up":
								return color.New(color.FgGreen).Sprint(sv.State)
							case "down":
								return color.New(color.FgRed).Sprint(sv.State)
							}
							return ""
						}(sv),
						func(sv inits.Service) string {
							switch sv.Enabled {
							case true:
								return color.New(color.FgGreen).Sprint(sv.Enabled)
							case false:
								return color.New(color.FgRed).Sprint(sv.Enabled)
							}
							return ""
						}(sv),
						func(sv inits.Service) string {
							switch sv.State {
							case "up":
								return color.New(color.FgHiMagenta).Sprint(sv.PID)
							case "down":
								return color.New(color.FgHiBlack).Sprint("---")
							}
							return ""
						}(sv),
						strings.Join(sv.Command, " "),
					)
				}

			}
		},
	}
)

func init() {
	cmd_root.AddCommand(
		cmd_query_list,
		cmd_query_list_available,
		cmd_query_status,
	)

	cmd_query_status.Flags().BoolVar(&scriptable, "scriptable", false, "format output as csv")
}
