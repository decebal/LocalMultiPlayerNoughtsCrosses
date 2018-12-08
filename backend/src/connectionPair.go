package main

import (
	"log"
	"sync"
	"time"
)

const timeoutBeforeReBroadcast = 5
const timeoutBeforeConnectionDrop = 1

type connectionPair struct {
	connectionsMx sync.RWMutex
	connections   map[*connection]struct{}
	receiveMove   chan bool
	logMx         sync.RWMutex
	log           [][]byte
	gs            gameState
}

func newConnectionPair() *connectionPair {
	cp := &connectionPair{
		connectionsMx: sync.RWMutex{},
		receiveMove:   make(chan bool),
		connections:   make(map[*connection]struct{}),
		gs:            newGameState(),
	}

	go func() {
		for {
			select {
				case <-cp.receiveMove:
				case <-time.After(timeoutBeforeReBroadcast * time.Second):
			}

			cp.connectionsMx.RLock()
			for c := range cp.connections {
				select {
				case c.doBroadcast <- true:
				case <-time.After(timeoutBeforeConnectionDrop * time.Second):
					cp.removeConnection(c)
				}
			}
			cp.connectionsMx.RUnlock()
		}
	}()
	return cp
}

func (h *connectionPair) addConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	h.connections[conn] = struct{}{}

}

func (h *connectionPair) removeConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.doBroadcast)
	}
	log.Println("Player disconnected")
	h.gs.resetGame()
}
