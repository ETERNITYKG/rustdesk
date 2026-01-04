package server

import (
	"context"
	"sync"
)

// AudioService models the audio capture pipeline (audio_service.rs).
// Data flow: platform driver -> AudioService.ReadFrame -> connection stream.
type AudioService interface {
	Start(ctx context.Context, cfg AudioConfig) error
	Stop() error
	SetInputDevice(name string) error
	ReadFrame(ctx context.Context) (AudioFrame, error)
}

type audioService struct {
	driver AudioPlatformDriver
	mu     sync.Mutex
	cfg    AudioConfig
}

func NewAudioService() AudioService {
	return &audioService{driver: newAudioPlatformDriver()}
}

func (svc *audioService) Start(ctx context.Context, cfg AudioConfig) error {
	svc.mu.Lock()
	svc.cfg = cfg
	svc.mu.Unlock()
	return svc.driver.Start(ctx, cfg)
}

func (svc *audioService) Stop() error {
	return svc.driver.Stop()
}

func (svc *audioService) SetInputDevice(name string) error {
	svc.mu.Lock()
	svc.cfg.InputDevice = name
	svc.mu.Unlock()
	return svc.driver.SetInputDevice(name)
}

func (svc *audioService) ReadFrame(ctx context.Context) (AudioFrame, error) {
	return svc.driver.ReadFrame(ctx)
}
