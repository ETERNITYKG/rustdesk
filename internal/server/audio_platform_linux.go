//go:build linux

package server

import (
	"context"
	"fmt"
)

type linuxAudioDriver struct{}

func newAudioPlatformDriver() AudioPlatformDriver {
	return &linuxAudioDriver{}
}

func (driver *linuxAudioDriver) Start(_ context.Context, _ AudioConfig) error {
	return fmt.Errorf("audio capture on Linux: %w", ErrNotImplemented)
}

func (driver *linuxAudioDriver) Stop() error {
	return nil
}

func (driver *linuxAudioDriver) SetInputDevice(_ string) error {
	return nil
}

func (driver *linuxAudioDriver) ReadFrame(_ context.Context) (AudioFrame, error) {
	return AudioFrame{}, fmt.Errorf("audio capture on Linux: %w", ErrNotImplemented)
}
