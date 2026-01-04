package server

import "context"

// InputService models the input injection pipeline (input_service.rs).
// Data flow: connection -> InputService.Dispatch -> platform driver.
type InputService interface {
	Start(ctx context.Context) error
	Stop() error
	Dispatch(ctx context.Context, event InputEvent) error
}

type inputService struct {
	driver InputPlatformDriver
}

func NewInputService() InputService {
	return &inputService{driver: newInputPlatformDriver()}
}

func (svc *inputService) Start(ctx context.Context) error {
	return svc.driver.Start(ctx)
}

func (svc *inputService) Stop() error {
	return svc.driver.Stop()
}

func (svc *inputService) Dispatch(ctx context.Context, event InputEvent) error {
	return svc.driver.Dispatch(ctx, event)
}
