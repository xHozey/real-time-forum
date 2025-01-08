package handlers

import (
	"net/http"
	"strconv"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) PostHandler(w http.ResponseWriter, r *http.Request) {
	post := types.Post{}
	utils.DecodeRequest(r, &post)
	err := db.HandlerDB.ValidatePost(post)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err.Error())
		return
	}
	_, post.User = db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	err = db.HandlerDB.ServiceDB.MiddlewareData.InsertPost(post)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, "Internal server error")
		return
	}
}

func (db *HandlerLayer) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
}

func (db *HandlerLayer) GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
	}
}
