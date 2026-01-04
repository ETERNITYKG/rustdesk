package server

import "context"

// VideoPlatformDriver abstracts platform-specific capture implementations.
type VideoPlatformDriver interface {
	Start(ctx context.Context, cfg VideoConfig) error
	Stop() error
	ReadFrame(ctx context.Context) (VideoFrame, error)
}
