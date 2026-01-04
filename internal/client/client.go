package client

// Rust mapping: src/client.rs and src/client/*

type MediaSender struct{}

type PointerEvent struct{}

type Interface interface{}

type Service interface {
	GetKeyState(key string) bool
	StartVideoThread(adapter any, onFrame any) error
	StartAudioThread() MediaSender
	SendMouse(event any) error
	SendPointerDeviceEvent(event PointerEvent) error
	InputOSPassword(password string, activate bool, ui Interface)
	HandleLoginError(code string, title string, text string, retryForRelay bool) bool
	CheckIfRetry(msgType string, title string, text string, retryForRelay bool) bool
}
