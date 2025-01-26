package websocket

import (
	"log"
	"strings"
	"time"
)

func (c *Client) read() {
	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}
		if len(msg.Content) > 2000 || len(strings.TrimSpace(msg.Content)) == 0 {
			continue
		}
		if msg.SocketType == "typing" {
			handleTyping(msg, c.Id)
			continue
		}
		c.Db.handleMsg(msg, c.Id)
	}
}

func handleTyping(msg Message, id int) {
	client, exist := Clients[msg.Target]
	message := Message{SocketType: "typing", Sender: id}
	if msg.Content == "stop-typing" {
		message.SocketType = "stop-typing"
	}
	if exist {
		client.Conn.WriteJSON(&message)
	}
}

func (db *WSlayer) handleMsg(msg Message, id int) {
	message := Message{SocketType: "chat", Sender: id, Content: msg.Content, Creation: time.Now()}
	client, exist := Clients[msg.Target]
	if exist {
		err := db.Data.InsertUserMessages(id, msg.Target, msg.Content)
		if err != nil {
			log.Println(err)
			return
		}
		client.Conn.WriteJSON(&message)
	}
}

func (c *Client) notify() {
	for _, client := range Clients {
		if client.Id != c.Id {
			c.SocketType = "status"
			err := client.Conn.WriteJSON(&c)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
