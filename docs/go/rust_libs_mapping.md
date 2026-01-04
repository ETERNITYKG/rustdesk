# Rust libs to Go capability mapping

This document maps Rust `libs/*` capabilities to proposed Go dependencies or
first-party packages for the Go port.

| Rust lib | Capability | Go mapping |
| --- | --- | --- |
| `libs/clipboard` | Clipboard abstraction across platforms | `golang.design/x/clipboard`, plus `go/core/clipboard` wrapper |
| `libs/enigo` | Input injection (keyboard/mouse) | `github.com/go-vgo/robotgo` or OS-specific implementations under `go/core/input` |
| `libs/hbb_common` | Shared types, config, protocol helpers | `go/core/common` (own types), `google.golang.org/protobuf` for protocol buffers |
| `libs/portable` | Portable build logic & runtime layout | `go/core/portable` + scripts in `scripts/go/` |
| `libs/remote_printer` | Remote printing support | `go/core/printing` with OS print APIs (`golang.org/x/sys` or platform SDKs) |
| `libs/scrap` | Screen capture | `github.com/kbinani/screenshot` or platform-specific capture packages in `go/core/capture` |
| `libs/virtual_display` | Virtual display management | `go/core/display` with Windows/Mac/Linux SDK integrations |

Notes:
- These are placeholders to anchor the Go module/package layout.
- Each Go package should document its platform support and required build tags.
