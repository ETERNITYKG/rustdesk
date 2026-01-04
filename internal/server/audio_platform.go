package server

import "context"

// AudioPlatformDriver abstracts platform-specific audio capture.
type AudioPlatformDriver interface {
	Start(ctx context.Context, cfg AudioConfig) error
	Stop() error
	SetInputDevice(name string) error
	ReadFrame(ctx context.Context) (AudioFrame, error)
}
