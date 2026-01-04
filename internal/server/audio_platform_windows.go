//go:build windows

package server

import (
	"context"
	"fmt"
)

type windowsAudioDriver struct{}

func newAudioPlatformDriver() AudioPlatformDriver {
	return &windowsAudioDriver{}
}

func (driver *windowsAudioDriver) Start(_ context.Context, _ AudioConfig) error {
	return fmt.Errorf("audio capture on Windows: %w", ErrNotImplemented)
}

func (driver *windowsAudioDriver) Stop() error {
	return nil
}

func (driver *windowsAudioDriver) SetInputDevice(_ string) error {
	return nil
}

func (driver *windowsAudioDriver) ReadFrame(_ context.Context) (AudioFrame, error) {
	return AudioFrame{}, fmt.Errorf("audio capture on Windows: %w", ErrNotImplemented)
}
