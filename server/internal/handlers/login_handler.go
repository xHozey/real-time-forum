package handlers

import (
	"net/http"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	userInfo := types.User{}
	err := utils.DecodeRequest(r, &userInfo)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, "Internal server error")
		
		return
	}
	userId, validation := db.HandlerDB.ValidateLogin(userInfo)
	if !validation {
		utils.SendResponseStatus(w, http.StatusBadRequest, "Invalid credential")
		return
	}
	err = db.HandlerDB.UpdateSession(w, userId)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, "Internal server error")
		return
	}
}
