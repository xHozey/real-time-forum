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
		handleMessage(msg)
	}
}

func handleMessage(message []byte) {
	parts := bytes.SplitN(message, []byte(" "), 2)
	if len(parts) != 2 {
		return
	}
	userid, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return
	}
	sendMessage(userid, parts[1])
}

func sendMessage(userId int, message []byte) {
	conn, exists := Clients[userId]
	if exists {
		conn.conn.WriteMessage(websocket.TextMessage, message)
	}
}
