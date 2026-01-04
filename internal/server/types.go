package server

import "time"

// Message represents a high-level transport payload flowing between services
// and the connection layer.
type Message struct {
	Kind string
	Data any
}

// AudioConfig models audio capture configuration aligned with Rust audio_service.
type AudioConfig struct {
	InputDevice string
	SampleRate  int
	Channels    int
}

// AudioFrame represents a single audio payload with timing data.
type AudioFrame struct {
	PCM       []float32
	Timestamp time.Time
}

// VideoSource matches the Rust video_service::VideoSource.
type VideoSource int

const (
	VideoSourceMonitor VideoSource = iota
	VideoSourceCamera
)

// VideoConfig models capture options for monitor/camera sources.
type VideoConfig struct {
	Source      VideoSource
	DisplayID   int
	RefreshRate int
}

// VideoFrame represents a captured video frame.
type VideoFrame struct {
	Bytes     []byte
	Width     int
	Height    int
	Format    string
	Timestamp time.Time
}

// InputEventType identifies the kind of input data being applied.
type InputEventType int

const (
	InputEventMouse InputEventType = iota
	InputEventKey
	InputEventPointer
	InputEventTouch
)

// MouseEvent mirrors pointer movement/click data flow.
type MouseEvent struct {
	X       int
	Y       int
	Button  string
	Pressed bool
	Wheel   int
}

// KeyEvent mirrors keyboard data flow.
type KeyEvent struct {
	KeyCode string
	Pressed bool
	Text    string
}

// PointerEvent mirrors richer pointer data flows (e.g. stylus).
type PointerEvent struct {
	X       int
	Y       int
	Buttons []string
}

// TouchEvent mirrors touch data flow.
type TouchEvent struct {
	X     int
	Y     int
	Scale float32
}

// InputEvent wraps all input payload variants.
type InputEvent struct {
	Type    InputEventType
	Mouse   *MouseEvent
	Key     *KeyEvent
	Pointer *PointerEvent
	Touch   *TouchEvent
}

// ClipboardFormat models clipboard payload kinds.
type ClipboardFormat string

const (
	ClipboardFormatText  ClipboardFormat = "text"
	ClipboardFormatImage ClipboardFormat = "image"
	ClipboardFormatFile  ClipboardFormat = "file"
)

// ClipboardData models clipboard transfer payloads.
type ClipboardData struct {
	Format ClipboardFormat
	Data   []byte
	Files  []string
	Width  int
	Height int
}
