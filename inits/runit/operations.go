package runit

import (
	"os"
	"os/exec"
	"path/filepath"

	"gitcat.ca/endigma/jasmine/util"
)

func (r *runit) Add(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	if !r.serviceEnabled(name) {
		if err := os.Symlink(filepath.Join(r.svdir, name), filepath.Join(r.runsvdir, name)); err != nil {
			return err
		}
	}

	return nil
}

func (r *runit) Remove(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	if r.serviceEnabled(name) {
		if err := os.Remove(filepath.Join(r.runsvdir, name)); err != nil {
			return err
		}
	}

	return nil
}

func (r *runit) Enable(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	if util.FileExists(filepath.Join(r.runsvdir, name, "down")) {
		if err := os.Remove(filepath.Join(r.runsvdir, name, "down")); err != nil {
			return err
		}
	}

	return nil
}

func (r *runit) Disable(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	if !util.FileExists(filepath.Join(r.runsvdir, name, "down")) {
		if _, err := os.Create(filepath.Join(r.runsvdir, name, "down")); err != nil {
			return err
		}
	}

	return nil
}

func (r *runit) Start(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	if err := r.serviceControl(name, controlUp); err != nil {
		return err
	}
	return nil
}

func (r *runit) Stop(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	if err := r.serviceControl(name, controlDown); err != nil {
		return err
	}
	return nil
}

func (r *runit) Restart(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	r.Stop(name)
	r.Start(name)
	return nil
}

func (r *runit) Reload(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	if err := r.serviceControl(name, controlHangup); err != nil {
		return err
	}
	return nil
}

func (r *runit) Once(name string) error {
	if err := r.serviceExists(name); err != nil {
		return err
	}

	if err := r.serviceControl(name, controlOnce); err != nil {
		return err
	}
	return nil
}

func (r *runit) Pass(cmd ...string) error {

	if err := exec.Command("sv", cmd...).Run(); err != nil {
		return err
	}
	return nil
}
