package websocket

import (
	"database/sql"

	"github.com/gorilla/websocket"
)

type WSlayer struct {
	Data *sql.DB
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type room struct {
	clients map[*client]bool
	forward chan []byte
	
}

type client struct {
	nickname string
	socket   *websocket.Conn
	receive  chan []byte
	room     *room
}
