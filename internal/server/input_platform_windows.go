//go:build windows

package server

import (
	"context"
	"fmt"
)

type windowsInputDriver struct{}

func newInputPlatformDriver() InputPlatformDriver {
	return &windowsInputDriver{}
}

func (driver *windowsInputDriver) Start(_ context.Context) error {
	return fmt.Errorf("input injection on Windows: %w", ErrNotImplemented)
}

func (driver *windowsInputDriver) Stop() error {
	return nil
}

func (driver *windowsInputDriver) Dispatch(_ context.Context, _ InputEvent) error {
	return fmt.Errorf("input injection on Windows: %w", ErrNotImplemented)
}
