package handlers

import (
	"fmt"
	"net/http"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) InfoHandler(w http.ResponseWriter, r *http.Request) {
	id, nickname := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	if id == 0 {
		utils.SendResponseStatus(w, http.StatusUnauthorized, types.ErrUnauthorized)
		return
	}
	users, err := db.HandlerDB.ServiceDB.MiddlewareData.GetAllUsers(id)
	if err != nil {
		fmt.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
		return
	}
	data := types.InfoUser{UserId: id, Nickname: nickname, Clients: users}

	utils.SendJsonData(w, &data)
}
