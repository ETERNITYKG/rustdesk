package server

import (
	"context"
)

// VideoService models the capture pipeline (video_service.rs).
// Data flow: platform driver -> VideoService.ReadFrame -> connection stream.
type VideoService interface {
	Start(ctx context.Context, cfg VideoConfig) error
	Stop() error
	ReadFrame(ctx context.Context) (VideoFrame, error)
}

type videoService struct {
	driver VideoPlatformDriver
}

func NewVideoService() VideoService {
	return &videoService{driver: newVideoPlatformDriver()}
}

func (svc *videoService) Start(ctx context.Context, cfg VideoConfig) error {
	return svc.driver.Start(ctx, cfg)
}

func (svc *videoService) Stop() error {
	return svc.driver.Stop()
}

func (svc *videoService) ReadFrame(ctx context.Context) (VideoFrame, error) {
	return svc.driver.ReadFrame(ctx)
}
