package server

import "context"

// ClipboardPlatformDriver abstracts platform clipboard access.
type ClipboardPlatformDriver interface {
	Start(ctx context.Context) error
	Stop() error
	Send(ctx context.Context, data ClipboardData) error
	Receive(ctx context.Context) (ClipboardData, error)
}
