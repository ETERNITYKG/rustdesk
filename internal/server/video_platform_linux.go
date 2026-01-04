//go:build linux

package server

import (
	"context"
	"fmt"
)

type linuxVideoDriver struct{}

func newVideoPlatformDriver() VideoPlatformDriver {
	return &linuxVideoDriver{}
}

func (driver *linuxVideoDriver) Start(_ context.Context, _ VideoConfig) error {
	return fmt.Errorf("video capture on Linux: %w", ErrNotImplemented)
}

func (driver *linuxVideoDriver) Stop() error {
	return nil
}

func (driver *linuxVideoDriver) ReadFrame(_ context.Context) (VideoFrame, error) {
	return VideoFrame{}, fmt.Errorf("video capture on Linux: %w", ErrNotImplemented)
}
