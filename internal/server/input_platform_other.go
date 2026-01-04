//go:build !windows && !linux && !darwin

package server

import (
	"context"
	"fmt"
)

type otherInputDriver struct{}

func newInputPlatformDriver() InputPlatformDriver {
	return &otherInputDriver{}
}

func (driver *otherInputDriver) Start(_ context.Context) error {
	return fmt.Errorf("input injection: %w", ErrNotImplemented)
}

func (driver *otherInputDriver) Stop() error {
	return nil
}

func (driver *otherInputDriver) Dispatch(_ context.Context, _ InputEvent) error {
	return fmt.Errorf("input injection: %w", ErrNotImplemented)
}
