package handlers

import (
	"log"
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
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
	data := types.InfoUser{UserId: id, Nickname: nickname, Clients: users}

	err = utils.SendJsonData(w, &data)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}
