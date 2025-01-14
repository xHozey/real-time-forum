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
	Clients = make(map[int]*Client)
	mu      sync.Mutex
)

type Client struct {
	Id     int `json:"id"`
	Conn   *websocket.Conn
	Status bool `json:"status"`
	Db     WSlayer
}

type Messages struct {
	
}