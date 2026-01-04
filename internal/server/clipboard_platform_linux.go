//go:build linux

package server

import (
	"context"
	"fmt"
)

type linuxClipboardDriver struct{}

func newClipboardPlatformDriver() ClipboardPlatformDriver {
	return &linuxClipboardDriver{}
}

func (driver *linuxClipboardDriver) Start(_ context.Context) error {
	return fmt.Errorf("clipboard sync on Linux: %w", ErrNotImplemented)
}

func (driver *linuxClipboardDriver) Stop() error {
	return nil
}

func (driver *linuxClipboardDriver) Send(_ context.Context, _ ClipboardData) error {
	return fmt.Errorf("clipboard sync on Linux: %w", ErrNotImplemented)
}

func (driver *linuxClipboardDriver) Receive(_ context.Context) (ClipboardData, error) {
	return ClipboardData{}, fmt.Errorf("clipboard sync on Linux: %w", ErrNotImplemented)
}
