package clipboard

// Rust mapping: src/clipboard.rs

type ClipboardSide int

type Clipboard struct{}

type MultiClipboards struct{}

type Message struct{}

type Service interface {
	CheckClipboard(side ClipboardSide, connID int, server bool) error
	IsFileURLSetByRustDesk(urls []string) bool
	CheckClipboardFiles(side ClipboardSide, connID int) error
	UpdateClipboardFiles(files []string, side ClipboardSide)
	TryEmptyClipboardFiles(side ClipboardSide, connID int)
	CheckClipboardCM() (MultiClipboards, error)
	UpdateClipboard(clipboards []Clipboard, side ClipboardSide)
	IsSupportMultiClipboard(peerVersion string, peerPlatform string) bool
	GetCurrentClipboardMessage(client bool) (Message, bool)
	HandleMsgClipboard(clipboard Clipboard)
	HandleMsgMultiClipboards(clipboards MultiClipboards)
	GetClipboardsMessage(client bool) (Message, bool)
}
