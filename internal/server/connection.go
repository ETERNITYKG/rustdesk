package server

import (
	"context"
	"sync"
)

// Connection abstracts the server connection logic (connection.rs).
// Data flow: network -> Connection.Receive -> service dispatch (input/clipboard) -> platform.
//
//	service capture (audio/video) -> Connection.Send -> network.
type Connection interface {
	ID() int
	Send(ctx context.Context, msg Message) error
	Receive(ctx context.Context) (Message, error)
	Close() error
}

// ConnectionManager tracks active connections for the server layer.
type ConnectionManager interface {
	Add(conn Connection)
	Remove(connID int)
	List() []Connection
}

type connectionManager struct {
	mu    sync.RWMutex
	conns map[int]Connection
}

func NewConnectionManager() ConnectionManager {
	return &connectionManager{conns: make(map[int]Connection)}
}

func (mgr *connectionManager) Add(conn Connection) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.conns[conn.ID()] = conn
}

func (mgr *connectionManager) Remove(connID int) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	delete(mgr.conns, connID)
}

func (mgr *connectionManager) List() []Connection {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	out := make([]Connection, 0, len(mgr.conns))
	for _, conn := range mgr.conns {
		out = append(out, conn)
	}
	return out
}
