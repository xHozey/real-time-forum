package websocket

import (
	"sync"

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

var (
	Clients = make(map[int]*client)
	mu      sync.Mutex
)

type client struct {
	id     int
	conn   *websocket.Conn
	status bool
	db     WSlayer
}
