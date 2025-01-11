package websocket

import (
	"forum/server/internal/data"

	"github.com/gorilla/websocket"
)

type WSlayer struct {
	Data data.DataLayer
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Clients = make(map[int]*client)

type client struct {
	id int
	conn *websocket.Conn
	send chan []byte
	status bool
}
