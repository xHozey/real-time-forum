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
		return
	}
	db.Data.DataDB.Exec("UPDATE user_profile SET status = 1 WHERE id = ?", id)
	client := &Client{Id: id, Status: true, Conn: conn, Db: *db, Nickname: nickname}
	Mu.Lock()
	_, exists := Clients[id]
	if !exists {
		Clients[id] = client
	}
	Clients[id].Window++
	Mu.Unlock()
	go client.notify()
	client.read()
	defer client.CloseConn(db.Data.DataDB, id, false)
}
