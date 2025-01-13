package handlers

import (
	"net/http"

	"forum/server/internal/utils"
)

func (db *HandlerLayer) InfoHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	users, err := db.HandlerDB.ServiceDB.MiddlewareData.GetAllUsers(id)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
	}
	utils.SendJsonData(w, &users)
}
