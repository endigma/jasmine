package runit

import (
	"gitcat.ca/endigma/jasmine/inits"
	"github.com/spf13/viper"
)

type runit struct {
	svdir    string
	runsvdir string
}

// New creates a new runit init object
func New() inits.Init {
	return &runit{
		svdir:    viper.GetString("runit.svdir"),
		runsvdir: viper.GetString("runit.runsvdir"),
	}
}

type control []byte

var (
	controlUp     control = []byte("u")
	controlDown   control = []byte("d")
	controlOnce   control = []byte("o")
	controlHangup control = []byte("h")
)
