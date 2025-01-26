package handlers

import (
	"log"
	"net/http"
	"strings"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	userInfo := types.User{}
	err := utils.DecodeRequest(r, &userInfo)
	userInfo.Nickname = strings.ToLower(userInfo.Nickname)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusBadRequest, types.ErrBadRequest)
		return
	}
	userId, err := db.HandlerDB.ValidateLogin(userInfo)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	err = db.HandlerDB.UpdateSession(w, userId)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}
