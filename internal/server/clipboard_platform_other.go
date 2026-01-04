//go:build !windows && !linux && !darwin

package server

import (
	"context"
	"fmt"
)

type otherClipboardDriver struct{}

func newClipboardPlatformDriver() ClipboardPlatformDriver {
	return &otherClipboardDriver{}
}

func (driver *otherClipboardDriver) Start(_ context.Context) error {
	return fmt.Errorf("clipboard sync: %w", ErrNotImplemented)
}

func (driver *otherClipboardDriver) Stop() error {
	return nil
}

func (driver *otherClipboardDriver) Send(_ context.Context, _ ClipboardData) error {
	return fmt.Errorf("clipboard sync: %w", ErrNotImplemented)
}

func (driver *otherClipboardDriver) Receive(_ context.Context) (ClipboardData, error) {
	return ClipboardData{}, fmt.Errorf("clipboard sync: %w", ErrNotImplemented)
}
