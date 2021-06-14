package inits

import (
	"fmt"
	"time"
)

type Init interface {
	// Enable provides a method to enable a service by name
	Add(name string) error

	// Disable provides a method to disable a service by name
	Remove(name string) error

	// Enable provides a method to up a service by name
	Enable(name string) error

	// Disable provides a method to down a service by name
	Disable(name string) error

	// Start provides a method to start a service by name
	Start(name string) error

	// Stop provides a method to stop a service by name
	Stop(name string) error

	// Restart provides a method to restart a service by name
	Restart(name string) error

	// Reload provides a method to reload a service by name
	Reload(name string) error

	// Once provides a method to run a service once by name
	Once(name string) error

	// Pass allows sending commands directly to an init system's official control system
	Pass(cmd ...string) error

	// List provides a method to list enabled services, optionally filtered by names
	List([]string) ([]Service, error)

	// List provides a method to disable a service by name
	ListAvailable() (map[string]bool, error)

	// Status provides a method to view the status of a single service by name
	Status(service string) (Service, error)
}

type Service struct {
	Name    string
	State   string
	Enabled bool
	PID     int64
	Command []string
	Uptime  time.Duration
}

var (
	ErrServiceNotFound = func(sv string) error {
		return fmt.Errorf("Service '%s' not found", sv)
	}
	ErrServiceMalformed = func(sv string) error {
		return fmt.Errorf("Service '%s' is malformed", sv)
	}
	ErrServiceAlreadyEnabled = func(sv string) error {
		return fmt.Errorf("Service '%s' is already enabled", sv)
	}
	ErrServiceAlreadyDisabled = func(sv string) error {
		return fmt.Errorf("Service '%s' is already disabled", sv)
	}
	ErrServiceAlreadyUp = func(sv string) error {
		return fmt.Errorf("Service '%s' is already up", sv)
	}
	ErrServiceAlreadyDown = func(sv string) error {
		return fmt.Errorf("Service '%s' is already down", sv)
	}
)
