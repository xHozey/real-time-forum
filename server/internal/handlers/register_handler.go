package handlers

import (
	"log"
	"net/http"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	userInfo := types.User{}
	err := utils.DecodeRequest(r, &userInfo)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusBadRequest, types.ErrBadRequest)
		return
	}
	if err := db.HandlerDB.ValidateCredentials(userInfo); err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	if err := db.HandlerDB.InsertClearCredential(userInfo); err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}
