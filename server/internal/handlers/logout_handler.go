package handlers

import (
	"net/http"

	"forum/server/internal/utils"
	"forum/server/internal/websocket"
)

func (db *HandlerLayer) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	if id != 0 {
		db.HandlerDB.ServiceDB.MiddlewareData.DeleteSession(id)
		utils.DeleteCookie(w)
		client, exist := websocket.Clients[id]
		if exist {
			
			client.CloseConn(db.HandlerDB.ServiceDB.MiddlewareData.DataDB, id, true)
		}
	}
}
