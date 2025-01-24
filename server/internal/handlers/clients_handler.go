package handlers

import (
	"log"
	"net/http"
	"strconv"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) ClientHandler(w http.ResponseWriter, r *http.Request) {
	source, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	target, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusNotFound, types.ErrNotExist)
		return
	}
	offset, err := utils.GetOffset(w, r)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
	messages, err := db.HandlerDB.ServiceDB.MiddlewareData.GetConverceation(source, target, offset)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
	err = utils.SendJsonData(w, messages)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}
