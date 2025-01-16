package websocket

import (
	"bytes"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

func (c *Client) read() {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		c.Db.handleMsg(msg, c.Id)
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
		err := db.Data.InsertUserMessages(id, target, string(parts[1]))
		if err != nil {
			log.Println(err)
			return
		}
		client.Conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (c *Client) notify() {
	for _, client := range Clients {
		if client.Id != c.Id {
			client.Conn.WriteJSON(&c)
		}
	}
}
