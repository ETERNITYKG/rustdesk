//go:build linux

package server

import (
	"context"
	"fmt"
)

type linuxInputDriver struct{}

func newInputPlatformDriver() InputPlatformDriver {
	return &linuxInputDriver{}
}

func (driver *linuxInputDriver) Start(_ context.Context) error {
	return fmt.Errorf("input injection on Linux: %w", ErrNotImplemented)
}

func (driver *linuxInputDriver) Stop() error {
	return nil
}

func (driver *linuxInputDriver) Dispatch(_ context.Context, _ InputEvent) error {
	return fmt.Errorf("input injection on Linux: %w", ErrNotImplemented)
}
