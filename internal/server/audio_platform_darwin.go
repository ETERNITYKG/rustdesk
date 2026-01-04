//go:build darwin

package server

import (
	"context"
	"fmt"
)

type darwinAudioDriver struct{}

func newAudioPlatformDriver() AudioPlatformDriver {
	return &darwinAudioDriver{}
}

func (driver *darwinAudioDriver) Start(_ context.Context, _ AudioConfig) error {
	return fmt.Errorf("audio capture on macOS: %w", ErrNotImplemented)
}

func (driver *darwinAudioDriver) Stop() error {
	return nil
}

func (driver *darwinAudioDriver) SetInputDevice(_ string) error {
	return nil
}

func (driver *darwinAudioDriver) ReadFrame(_ context.Context) (AudioFrame, error) {
	return AudioFrame{}, fmt.Errorf("audio capture on macOS: %w", ErrNotImplemented)
}
