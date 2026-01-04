package flutter

// Rust mapping: src/flutter.rs and src/flutter_ffi.rs

type SessionID string

type FileEntry struct{}

type Message struct{}

type StreamSink[T any] interface{}

type Session struct{}

type SessionManager interface {
	SessionAddExisted(id SessionID, peerID string) error
	SessionAdd(id SessionID, peerID string) error
	SessionStart(sessionID SessionID) error
	UpdateTextClipboardRequired()
	UpdateFileClipboardRequired()
	SendClipboardMessage(msg Message, isFile bool)
	MakeFDFlutter(id int32, entries []FileEntry, onlyCount bool) string
	GetCurSessionID() SessionID
	GetCurPeerID() string
	SetCurSessionID(sessionID SessionID)
	SessionGetRGBASize(sessionID SessionID, display int) int
	SessionNextRGBA(sessionID SessionID, display int)
	SessionSetSize(sessionID SessionID, display int, width int, height int)
	SessionRegisterPixelbufferTexture(sessionID SessionID, display int, ptr uintptr)
	PushSessionEvent(sessionID SessionID, name string, event map[string]string)
	PushGlobalEvent(channel string, event string) (bool, error)
	GetGlobalEventChannels() []string
	StartGlobalEventStream(sink StreamSink[string], appType string) error
	StopGlobalEventStream(appType string)
	SessionSendPointer(sessionID SessionID, msg string)
	GetCurSession() *Session
}
