package websocket

import (
	"log"
)

func (c *Client) read() {
	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}
		c.Db.handleMsg(msg, c.Id)
	}
}

func (db *WSlayer) handleMsg(msg Message, id int) {
	message := Message{SocketType: "chat", Sender: id, Content: msg.Content}
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
			client.Conn.WriteJSON(&c)
		}
	}
}
