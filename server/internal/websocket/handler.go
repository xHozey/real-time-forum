package websocket

import (
	"log"
	"net/http"
)

func (db *WSlayer) WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Upgrade fail", http.StatusUpgradeRequired)
	}
}
