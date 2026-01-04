//go:build windows

package server

import (
	"context"
	"fmt"
)

type windowsVideoDriver struct{}

func newVideoPlatformDriver() VideoPlatformDriver {
	return &windowsVideoDriver{}
}

func (driver *windowsVideoDriver) Start(_ context.Context, _ VideoConfig) error {
	return fmt.Errorf("video capture on Windows: %w", ErrNotImplemented)
}

func (driver *windowsVideoDriver) Stop() error {
	return nil
}

func (driver *windowsVideoDriver) ReadFrame(_ context.Context) (VideoFrame, error) {
	return VideoFrame{}, fmt.Errorf("video capture on Windows: %w", ErrNotImplemented)
}
