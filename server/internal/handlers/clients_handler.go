package handlers

import (
	"fmt"
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
	offset, err := utils.GetOffset(w,r)
	if err != nil {
		return
	}
	messages, err := db.HandlerDB.ServiceDB.MiddlewareData.GetConverceation(source, target, offset)
	if err != nil {
		fmt.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
		return
	}
	utils.SendJsonData(w, messages)
}

