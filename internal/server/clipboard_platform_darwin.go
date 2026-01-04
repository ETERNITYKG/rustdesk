//go:build darwin

package server

import (
	"context"
	"fmt"
)

type darwinClipboardDriver struct{}

func newClipboardPlatformDriver() ClipboardPlatformDriver {
	return &darwinClipboardDriver{}
}

func (driver *darwinClipboardDriver) Start(_ context.Context) error {
	return fmt.Errorf("clipboard sync on macOS: %w", ErrNotImplemented)
}

func (driver *darwinClipboardDriver) Stop() error {
	return nil
}

func (driver *darwinClipboardDriver) Send(_ context.Context, _ ClipboardData) error {
	return fmt.Errorf("clipboard sync on macOS: %w", ErrNotImplemented)
}

func (driver *darwinClipboardDriver) Receive(_ context.Context) (ClipboardData, error) {
	return ClipboardData{}, fmt.Errorf("clipboard sync on macOS: %w", ErrNotImplemented)
}
