package runit

import (
	"os"
	"os/exec"
	"path/filepath"

	"gitcat.ca/endigma/jasmine/util"
)

func (r *runit) Enable(services []string) error {
	if err := r.servicesExist(services); err != nil {
		return err
	}

	for _, sv := range services {
		if !r.serviceEnabled(sv) {
			if err := os.Symlink(filepath.Join(r.svdir, sv), filepath.Join(r.runsvdir, sv)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *runit) Disable(services []string) error {
	if err := r.servicesExist(services); err != nil {
		return err
	}

	for _, sv := range services {
		if r.serviceEnabled(sv) {
			if err := os.Remove(filepath.Join(r.runsvdir, sv)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *runit) Up(services []string) error {
	for _, sv := range services {
		if util.FileExists(filepath.Join(r.runsvdir, sv, "down")) {
			if err := os.Remove(filepath.Join(r.runsvdir, sv, "down")); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *runit) Down(services []string) error {
	for _, sv := range services {
		if !util.FileExists(filepath.Join(r.runsvdir, sv, "down")) {
			if _, err := os.Create(filepath.Join(r.runsvdir, sv, "down")); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *runit) Start(services []string) error {
	for _, sv := range services {
		if err := r.serviceControl(sv, controlUp); err != nil {
			return err
		}
	}
	return nil
}

func (r *runit) Stop(services []string) error {
	for _, sv := range services {
		if err := r.serviceControl(sv, controlDown); err != nil {
			return err
		}
	}
	return nil
}

func (r *runit) Restart(services []string) error {
	r.Stop(services)
	r.Start(services)
	return nil
}

func (r *runit) Reload(services []string) error {
	for _, sv := range services {
		if err := r.serviceControl(sv, controlHangup); err != nil {
			return err
		}
	}
	return nil
}

func (r *runit) Once(services []string) error {
	for _, sv := range services {
		if err := r.serviceControl(sv, controlOnce); err != nil {
			return err
		}
	}
	return nil
}

func (r *runit) Pass(cmd ...string) error {
	if err := exec.Command("sv", cmd...).Run(); err != nil {
		return err
	}
	return nil
}
