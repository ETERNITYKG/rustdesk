//go:build !windows && !linux && !darwin

package server

import (
	"context"
	"fmt"
)

type otherVideoDriver struct{}

func newVideoPlatformDriver() VideoPlatformDriver {
	return &otherVideoDriver{}
}

func (driver *otherVideoDriver) Start(_ context.Context, _ VideoConfig) error {
	return fmt.Errorf("video capture: %w", ErrNotImplemented)
}

func (driver *otherVideoDriver) Stop() error {
	return nil
}

func (driver *otherVideoDriver) ReadFrame(_ context.Context) (VideoFrame, error) {
	return VideoFrame{}, fmt.Errorf("video capture: %w", ErrNotImplemented)
}
