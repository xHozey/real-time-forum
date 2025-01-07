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
		err = utils.SendResponseStatus(w, http.StatusBadRequest, "Bad Request")
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	if err := db.HandlerDB.ValidateCredentials(userInfo); err != nil {
		err = utils.SendResponseStatus(w, http.StatusBadRequest, err.Error())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	if err := db.HandlerDB.InsertClearCredential(userInfo); err != nil {
		err = utils.SendResponseStatus(w, http.StatusInternalServerError, "Internal server error")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}
