package util

import (
	"fmt"

	"github.com/fatih/color"
)

func FancyWarn(str string, args ...interface{}) {
	color.New(color.FgRed).Print("\n[!] ")
	fmt.Printf(str, args...)
}
