package handlers

import (
	"log"
	"net/http"
	"strconv"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) PostHandler(w http.ResponseWriter, r *http.Request) {
	post := types.Post{}
	err := utils.DecodeRequest(r, &post)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusBadRequest, types.ErrBadRequest)
		return
	}
	id, nickname := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	if id == 0 {
		utils.SendResponseStatus(w, http.StatusUnauthorized, types.ErrUnauthorized)
		return
	}
	post.UserId = id
	err = db.HandlerDB.ValidatePost(post)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	insertedPost, err := db.HandlerDB.ServiceDB.MiddlewareData.InsertPost(post)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
	insertedPost.Author = nickname
	err = utils.SendJsonData(w, &insertedPost)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}

func (db *HandlerLayer) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	offset, err := utils.GetOffset(w, r)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
	posts, err := db.HandlerDB.ServiceDB.MiddlewareData.GetAllPosts(id, offset)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}

	if err := utils.SendJsonData(w, &posts); err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}

func (db *HandlerLayer) GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusNotFound, types.ErrNotExist)
		return
	}

	post, err := db.HandlerDB.ServiceDB.MiddlewareData.GetPostById(id)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}

	if err := utils.SendJsonData(w, &post); err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}
