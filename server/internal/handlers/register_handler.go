package handlers

import (
	"net/http"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	userInfo := types.User{}
	err := utils.DecodeRequest(r, &userInfo)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	if err := db.HandlerDB.ValidateCredentials(userInfo); err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	if err := db.HandlerDB.InsertClearCredential(userInfo); err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
		return
	}
}
