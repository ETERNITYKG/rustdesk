package portforward

// Rust mapping: src/port_forward.rs

type LoginConfig struct{}

type DataStream struct{}

type Interface interface{}

type Listener interface {
	Listen(id string, password string, port int, ui Interface, uiStream DataStream, key string, token string, cfg *LoginConfig, remoteHost string, remotePort int) error
}
