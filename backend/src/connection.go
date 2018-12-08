package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var connections []*connectionPair

type connection struct {
	doBroadcast chan bool
	cp *connectionPair
	playerNum int
}

type wsHandler struct{}

func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for {
		_, clientMoveMessage, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		field, _ := strconv.ParseInt(string(clientMoveMessage[:]), 10, 32)
		log.Printf("Player %v is making a move on %v", c.playerNum, int(field))
		c.cp.gs.makeMove(c.playerNum, int(field))
		c.cp.receiveMove <- true
	}
}

func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for range c.doBroadcast {
		sendPlayerInfoToConnection(wsConn, c)
		sendGameStateToConnection(wsConn, c)
	}
}

func getConnectionPairWithEmptySlot() (*connectionPair, int) {
	sizeBefore := len(connections)
	for _, h := range connections {
		if len(h.connections) == 1 {
			log.Printf("Players paired")
			return h, len(h.connections)
		}
	}

	h := newConnectionPair()
	connections = append(connections, h)
	log.Printf("Player seated in new connectionPair no. %v", len(connections))
	return connections[sizeBefore], 0
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading %s", err)
		return
	}

	cp, pn := getConnectionPairWithEmptySlot()
	c := &connection{doBroadcast: make(chan bool), cp: cp, playerNum: pn}
	c.cp.addConnection(c)

	if c.cp.gs.StatusMessage == resetWaitPaired {
		c.cp.gs = newGameState()
		c.cp.gs.numberOfPlayers = 1
		log.Println("game was reset")
	}

	c.cp.gs.addPlayer()
	c.cp.receiveMove <- true

	var wg sync.WaitGroup
	wg.Add(2)
	go c.writer(&wg, wsConn)
	go c.reader(&wg, wsConn)
	wg.Wait()
	wsConn.Close()
}

func sendGameStateToConnection(wsConn *websocket.Conn, c *connection) {
	err := wsConn.WriteMessage(websocket.TextMessage, c.cp.gs.gameStateToJSON())
	if err != nil {
		c.cp.removeConnection(c)
	}
}

func sendPlayerInfoToConnection(wsConn *websocket.Conn, c *connection) {
	playerInfo, _ := json.Marshal(map[string]int{"playerNum": c.playerNum})
	_ = wsConn.WriteMessage(websocket.TextMessage, playerInfo)
}
