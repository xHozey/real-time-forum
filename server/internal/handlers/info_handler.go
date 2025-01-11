package handlers

import (
	"net/http"

	"forum/server/internal/types"
	"forum/server/internal/utils"
	"forum/server/internal/websocket"
)

func (db *HandlerLayer) InfoHandler(w http.ResponseWriter, r *http.Request) {
	userInfo := []types.InfoUser{}
	id, user := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	for users := range websocket.Clients {
		if id != users {
			userInfo = append(userInfo, types.InfoUser{User_id: users, Nickname: user})
		}
	}
	utils.SendJsonData(w, &userInfo)
}
