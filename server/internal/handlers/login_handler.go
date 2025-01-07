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
		err = utils.SendResponseStatus(w, http.StatusInternalServerError, "Internal server error")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	userId, validation := db.HandlerDB.ValidateLogin(userInfo)
	if !validation {
		err = utils.SendResponseStatus(w, http.StatusBadRequest, "Invalid credential")
		if err != nil {
			http.Error(w, "Invalid credential", http.StatusBadRequest)
		}
		return
	}
	err = db.HandlerDB.UpdateSession(w, userId)
	if err != nil {
		err = utils.SendResponseStatus(w, http.StatusInternalServerError, "Internal server error")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
}
