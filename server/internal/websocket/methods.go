package websocket

import (
	"github.com/gorilla/websocket"
)

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		for client := range r.clients {
			msg := <-r.forward
			client.receive <- msg
		}
	}
}
