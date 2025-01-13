package websocket

import (
	"log"
	"net/http"

	"forum/server/internal/utils"
)

func (db *WSlayer) WsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := db.Data.GetUserBySession(utils.GetCookie(r))
	if id == 0 {
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Upgrade fail", http.StatusUpgradeRequired)
	}
	client := &client{id: id, status: true, conn: conn}
	mu.Lock()
	Clients[id] = client
	mu.Unlock()
	client.read()
	defer func() {
		client.status = false
		conn.Close()
	}()
}
