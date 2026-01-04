//go:build darwin

package server

import (
	"context"
	"fmt"
)

type darwinVideoDriver struct{}

func newVideoPlatformDriver() VideoPlatformDriver {
	return &darwinVideoDriver{}
}

func (driver *darwinVideoDriver) Start(_ context.Context, _ VideoConfig) error {
	return fmt.Errorf("video capture on macOS: %w", ErrNotImplemented)
}

func (driver *darwinVideoDriver) Stop() error {
	return nil
}

func (driver *darwinVideoDriver) ReadFrame(_ context.Context) (VideoFrame, error) {
	return VideoFrame{}, fmt.Errorf("video capture on macOS: %w", ErrNotImplemented)
}
