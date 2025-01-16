package handlers

import (
	"net/http"
	"strconv"

	"forum/server/internal/utils"
)

func (db *HandlerLayer) ClientHandler(w http.ResponseWriter, r *http.Request) {
	source, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	target, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendResponseStatus(w, http.StatusNotFound, err)
		return
	}
	messages, err := db.HandlerDB.ServiceDB.MiddlewareData.GetConverceation(source, target)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
		return
	}
	utils.SendJsonData(w, messages)
}
