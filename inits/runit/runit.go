package runit

import (
	"time"

	"gitcat.ca/endigma/jasmine/inits"
)

type runit struct {
	timeout  time.Duration
	svdir    string
	runsvdir string
}

// New creates a new runit init with the specified options
func New(svdir, runsvdir string, t int) inits.Init {
	return &runit{
		timeout:  time.Duration(t) * time.Second,
		svdir:    svdir,
		runsvdir: runsvdir,
	}
}

type control []byte

var (
	controlUp        control = []byte("u")
	controlDown      control = []byte("d")
	controlOnce      control = []byte("o")
	controlPause     control = []byte("p")
	controlContinue  control = []byte("c")
	controlHangup    control = []byte("h")
	controlAlarm     control = []byte("a")
	controlInterrupt control = []byte("i")
	controlQuit      control = []byte("q")
	controlUSR1      control = []byte("1")
	controlUSR2      control = []byte("2")
	controlTerminate control = []byte("t")
	controlKill      control = []byte("k")
	controlExit      control = []byte("e")
)
