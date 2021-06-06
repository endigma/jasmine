package util

import (
	"fmt"
	"os/user"

	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func FancyWarn(str string, args ...interface{}) {
	color.New(color.FgRed).Print("\n[!] ")
	fmt.Printf(str, args...)
}

func SudoWarn() {
	if !func() bool {
		currentUser, err := user.Current()
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to get current user")
		}
		return currentUser.Username == "root"
	}() {
		if !viper.GetBool("suppress_permissions_warning") {
			FancyWarn("Jasmine must be run as root for this command to work\n\n")
		}
	}
}
