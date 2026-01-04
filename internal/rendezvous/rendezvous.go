package rendezvous

// Rust mapping: src/rendezvous_mediator.rs

type ServerHandle struct{}

type Mediator interface {
	Restart()
	StartAll()
	Start(server *ServerHandle, host string) error
	StartUDP(server *ServerHandle, host string) error
}
