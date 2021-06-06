package runit

import (
	"os"

	"gitcat.ca/endigma/jasmine/inits"
	"gitcat.ca/endigma/jasmine/util"
)

func (r *runit) List(services []string) ([]inits.Service, error) {
	var servicefiles []inits.Service
	var list []string

	if len(services) > 0 {
		list = services
	} else {
		files, err := os.ReadDir(r.runsvdir)
		for _, f := range files {
			list = append(list, f.Name())
		}
		if err != nil {
			return []inits.Service{}, err
		}
	}

	for _, f := range list {
		sv, err := r.serviceScan(f)
		if err != nil {
			return []inits.Service{}, err
		}
		servicefiles = append(servicefiles, sv)
	}

	return servicefiles, nil
}

func (r *runit) ListAvailable() (map[string]bool, error) {
	var list map[string]bool = make(map[string]bool)

	files, err := os.ReadDir(r.svdir)
	for _, f := range files {
		list[f.Name()] = util.FileExists(r.runsvdir + "/" + f.Name())
	}
	if err != nil {
		return list, err
	}

	return list, nil
}

func (r *runit) Status(services []string) ([]inits.Service, error) {
	return []inits.Service{}, nil
}
