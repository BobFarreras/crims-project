package ws

import (
	"testing"
)

func TestHub_RegisterAndUnregister(t *testing.T) {
	hub := &Hub{
		clients:    map[string]map[*Client]struct{}{},
		register:   make(chan *Client, 1),
		unregister: make(chan *Client, 1),
		broadcast:  make(chan Message, 1),
	}
	client := &Client{gameID: "game-1", send: make(chan []byte, 1)}

	hub.register <- client
	hub.RunOnce()
	hub.unregister <- client
	hub.RunOnce()

	if len(hub.clients["game-1"]) != 0 {
		t.Fatalf("expected no clients, got %d", len(hub.clients["game-1"]))
	}
}
