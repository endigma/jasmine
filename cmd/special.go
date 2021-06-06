package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	cmd_version = &cobra.Command{
		Use:   "version",
		Short: "print version and debug information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(
				color.New(color.FgHiBlack).Sprint("["),
				color.New(color.FgMagenta).Sprint("Jasmine v0.0.1"),
				color.New(color.FgHiBlack).Sprint("]"),
				"\n",
				"source: https://gitcat.ca/endigma/jasmine\n")
		},
	}
)

func init() {
	cmd_root.AddCommand(cmd_version)
}
