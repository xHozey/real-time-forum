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
	SocketType string `json:"type"`
	Id         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Conn       *websocket.Conn
	Status     bool `json:"status"`
	Db         WSlayer
}

type Message struct {
	SocketType string `json:"type"`
	Sender     int    `json:"sender"`
	Target     int    `json:"target"`
	Content    string `json:"content"`
}
