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
	post.UserId, _ = db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	if post.UserId == 0 {
		utils.SendResponseStatus(w, http.StatusUnauthorized, types.ErrUnauthorized)
		return
	}
	err := db.HandlerDB.ValidatePost(post)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	err = db.HandlerDB.ServiceDB.MiddlewareData.InsertPost(post)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
		return
	}
}

func (db *HandlerLayer) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	posts, err := db.HandlerDB.ServiceDB.MiddlewareData.GetAllPosts(id)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.SendJsonData(w, &posts); err != nil {
		return
	}
}

func (db *HandlerLayer) GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendResponseStatus(w, http.StatusNotFound, err)
		return
	}

	post, err := db.HandlerDB.ServiceDB.MiddlewareData.GetPostById(id)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
	}

	if err := utils.SendJsonData(w, &post); err != nil {
		return
	}
}
