package server

import "context"

// ClipboardService models clipboard synchronization (clipboard_service.rs).
// Data flow: platform driver -> ClipboardService.Receive -> connection -> remote.
//
//	connection -> ClipboardService.Send -> platform driver.
type ClipboardService interface {
	Start(ctx context.Context) error
	Stop() error
	Send(ctx context.Context, data ClipboardData) error
	Receive(ctx context.Context) (ClipboardData, error)
}

type clipboardService struct {
	driver ClipboardPlatformDriver
}

func NewClipboardService() ClipboardService {
	return &clipboardService{driver: newClipboardPlatformDriver()}
}

func (svc *clipboardService) Start(ctx context.Context) error {
	return svc.driver.Start(ctx)
}

func (svc *clipboardService) Stop() error {
	return svc.driver.Stop()
}

func (svc *clipboardService) Send(ctx context.Context, data ClipboardData) error {
	return svc.driver.Send(ctx, data)
}

func (svc *clipboardService) Receive(ctx context.Context) (ClipboardData, error) {
	return svc.driver.Receive(ctx)
}
