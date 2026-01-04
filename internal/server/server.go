package server

// Rust mapping: src/server.rs and src/server/*

type Connection struct{}

type ServerHandle struct{}

type Manager interface {
	New() *ServerHandle
	CheckZombie()
}
