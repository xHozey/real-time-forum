package websocket

import (
	"log"
	"net/http"

	"forum/server/internal/utils"
)

func (db *WSlayer) WsHandler(w http.ResponseWriter, r *http.Request) {
	id, nickname := db.Data.GetUserBySession(utils.GetCookie(r))
	if id == 0 {
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Upgrade fail", http.StatusUpgradeRequired)
	}
	db.Data.DataDB.Exec("UPDATE user_profile SET status = 1 WHERE id = ?", id)
	client := &Client{Id: id, Status: true, Conn: conn, Db: *db, Nickname: nickname}
	mu.Lock()
	Clients[id] = client
	mu.Unlock()
	client.notify()
	client.read()
	defer func() {
		client.Status = false
		client.notify()
		db.Data.DataDB.Exec("UPDATE user_profile SET status = 0 WHERE id = ?", id)
		conn.Close()
	}()
}
