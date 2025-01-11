package websocket

import (
	"log"
	"net/http"

	"forum/server/internal/utils"
)

func (db *WSlayer) WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Upgrade fail", http.StatusUpgradeRequired)
	}

	id, _ := db.Data.GetUserBySession(utils.GetCookie(r))
	if id == 0 {
		conn.Close()
	}
	client := &client{conn: conn, id: id, send: make(chan []byte), status: true}
	Clients[id] = client
	client.read()
	defer func(){
		client.status = false
		conn.Close()
	}() 
}
