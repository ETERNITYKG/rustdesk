package common

// Rust mapping: src/common.rs (pub mod common in src/lib.rs)

type Service interface {
	GlobalInit() bool
	GlobalClean()
	CheckPort(host string, port int) string
	IncreasePort(host string, offset int) string
	GetAppName() string
	IsServer() bool
	IsCustomClient() bool
}
