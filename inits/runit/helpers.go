package runit

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gitcat.ca/endigma/jasmine/inits"
	"gitcat.ca/endigma/jasmine/util"
	"github.com/shirou/gopsutil/v3/process"
)

func (r *runit) scanService(name string) (inits.Service, error) {
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
	var command string

	if status == "up" {
		pidf, err := util.ReadFileContent(r.runsvdir + "/" + name + "/supervise/pid")
		if err != nil {
			return inits.Service{}, err
		}

		pid, err = strconv.Atoi(pidf)
		if err != nil {
			return inits.Service{}, err
		}

		ps, err := process.NewProcess(int32(pid))
		if err != nil {
			return inits.Service{}, err
		}

		commandslice, err := ps.CmdlineSlice()
		if err != nil {
			return inits.Service{}, err
		}
		command = commandslice[0]
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

func (r *runit) servicesExist(services []string) error {
	for _, sv := range services {
		if err := r.serviceExists(sv); err != nil {
			return err
		}
	}
	return nil
}

func (r *runit) serviceEnabled(name string) bool {
	return util.FileExists(r.runsvdir + "/" + name)
}
