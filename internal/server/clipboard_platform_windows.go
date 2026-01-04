//go:build windows

package server

import (
	"context"
	"fmt"
)

type windowsClipboardDriver struct{}

func newClipboardPlatformDriver() ClipboardPlatformDriver {
	return &windowsClipboardDriver{}
}

func (driver *windowsClipboardDriver) Start(_ context.Context) error {
	return fmt.Errorf("clipboard sync on Windows: %w", ErrNotImplemented)
}

func (driver *windowsClipboardDriver) Stop() error {
	return nil
}

func (driver *windowsClipboardDriver) Send(_ context.Context, _ ClipboardData) error {
	return fmt.Errorf("clipboard sync on Windows: %w", ErrNotImplemented)
}

func (driver *windowsClipboardDriver) Receive(_ context.Context) (ClipboardData, error) {
	return ClipboardData{}, fmt.Errorf("clipboard sync on Windows: %w", ErrNotImplemented)
}
