package platform

// Rust mapping: src/platform/* (re-exported in src/lib.rs via platform::*)
// Platform exposes core OS integration points that Rust re-exports at the crate root.

type CursorData struct{}

type DisplayInfo struct{}

type Service interface {
	GetCursor() (uint64, error)
	GetCursorData(handle uint64) (CursorData, error)
	GetCursorPos() (x int32, y int32, ok bool)
	GetFocusedDisplay(displays []DisplayInfo) (index int, ok bool)
	StartOSService() error
}
