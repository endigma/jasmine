package runit

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gitcat.ca/endigma/jasmine/inits"
	"gitcat.ca/endigma/jasmine/util"
)

func (r *runit) serviceScan(name string) (inits.Service, error) {
	if err := r.serviceExists(name); err != nil {
		return inits.Service{}, err
	}

	statusf, err := util.ReadFileContent(r.runsvdir + "/" + name + "/supervise/stat")
	if err != nil {
		return inits.Service{}, err
	}

	var status string
	switch statusf {
	case "run":
		status = "up"
	default:
		status = "down"
	}

	var pid int
	var command []string

	if status == "up" {
		pidf, err := util.ReadFileContent(r.runsvdir + "/" + name + "/supervise/pid")
		if err != nil {
			return inits.Service{}, err
		}

		pid, err = strconv.Atoi(pidf)
		if err != nil {
			return inits.Service{}, err
		}

		cmdline, err := ioutil.ReadFile(filepath.Join("/proc", pidf, "/cmdline"))
		if err != nil {
			return inits.Service{}, err
		}

		for _, a := range bytes.Split(cmdline, []byte{0}) {
			command = append(command, string(a))
		}
	}

	uptime, err := os.Stat(r.runsvdir + "/" + name + "/supervise/pid")
	if err != nil {
		return inits.Service{}, err
	}

	return inits.Service{
		Name:    name,
		State:   status,
		Enabled: !util.FileExists(filepath.Join(r.runsvdir, name, "down")),
		PID:     int64(pid),
		Command: command,
		Uptime:  time.Since(uptime.ModTime()),
	}, nil
}

func (r *runit) serviceControl(sv string, c control) error {
	file, err := os.OpenFile(filepath.Join(r.runsvdir, sv, "supervise", "control"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}

	file.Write(c)

	return nil
}

func (r *runit) serviceExists(name string) error {
	existsinsvdir := util.FileExists(r.svdir + "/" + name)
	existsinrunsvdir := util.FileExists(r.runsvdir + "/" + name)

	if existsinrunsvdir && !existsinsvdir {
		return inits.ErrServiceMalformed(name)
	}

	if !existsinsvdir {
		return inits.ErrServiceNotFound(name)
	}

	return nil
}

func (r *runit) serviceEnabled(name string) bool {
	return util.FileExists(r.runsvdir + "/" + name)
}
