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
		err = utils.SendResponseStatus(w, http.StatusBadRequest, "Bad Request")
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
}
