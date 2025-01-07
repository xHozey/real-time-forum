package handlers

import (
	"net/http"

	"forum/server/internal/utils"
)

func (db *HandlerLayer) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := db.HandlerDB.ServiceDB.GetUserBySession(utils.GetCookie(r))
	db.HandlerDB.ServiceDB.DeleteSession(id)
	utils.DeleteCookie(w)
}
