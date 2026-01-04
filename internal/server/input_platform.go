package server

import "context"

// InputPlatformDriver abstracts OS-specific input injection.
type InputPlatformDriver interface {
	Start(ctx context.Context) error
	Stop() error
	Dispatch(ctx context.Context, event InputEvent) error
}
