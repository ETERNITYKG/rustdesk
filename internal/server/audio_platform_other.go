//go:build !windows && !linux && !darwin

package server

import (
	"context"
	"fmt"
)

type otherAudioDriver struct{}

func newAudioPlatformDriver() AudioPlatformDriver {
	return &otherAudioDriver{}
}

func (driver *otherAudioDriver) Start(_ context.Context, _ AudioConfig) error {
	return fmt.Errorf("audio capture: %w", ErrNotImplemented)
}

func (driver *otherAudioDriver) Stop() error {
	return nil
}

func (driver *otherAudioDriver) SetInputDevice(_ string) error {
	return nil
}

func (driver *otherAudioDriver) ReadFrame(_ context.Context) (AudioFrame, error) {
	return AudioFrame{}, fmt.Errorf("audio capture: %w", ErrNotImplemented)
}
