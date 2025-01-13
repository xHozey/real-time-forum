package websocket

import (
	"bytes"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

func (c *client) read() {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		c.db.handleMsg(msg, c.id)
	}
}

func (db *WSlayer) handleMsg(msg []byte, id int) {
	parts := bytes.SplitN(msg, []byte(" "), 2)
	if len(parts) != 2 {
		return
	}
	target, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return
	}
	message := append([]byte(strconv.Itoa(id)), ' ')
	message = append(message, parts[1]...)
	client, exist := Clients[target]
	if exist {
		client.conn.WriteMessage(websocket.TextMessage, message)
	}
}
