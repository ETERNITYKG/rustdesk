//go:build darwin

package server

import (
	"context"
	"fmt"
)

type darwinInputDriver struct{}

func newInputPlatformDriver() InputPlatformDriver {
	return &darwinInputDriver{}
}

func (driver *darwinInputDriver) Start(_ context.Context) error {
	return fmt.Errorf("input injection on macOS: %w", ErrNotImplemented)
}

func (driver *darwinInputDriver) Stop() error {
	return nil
}

func (driver *darwinInputDriver) Dispatch(_ context.Context, _ InputEvent) error {
	return fmt.Errorf("input injection on macOS: %w", ErrNotImplemented)
}
