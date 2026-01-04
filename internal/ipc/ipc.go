package ipc

// Rust mapping: src/ipc.rs (pub mod ipc in src/lib.rs)

type API interface {
	UpdateTemporaryPassword() error
	GetPermanentPassword() string
	SetPermanentPassword(password string) error
	SetUnlockPin(pin string, translate bool) error
	GetUnlockPin() string
	GetTrustedDevices() string
	RemoveTrustedDevices(hwids [][]byte)
	ClearTrustedDevices()
	GetID() string
	SetOption(key string, value string)
	GetProxyStatus() bool
	CloseAllInstances() (bool, error)
}
