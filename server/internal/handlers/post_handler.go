package handlers

import (
	"fmt"
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
		utils.SendResponseStatus(w, http.StatusUnauthorized, types.ErrUnauthorized.Error())
		return
	}
	err := db.HandlerDB.ValidatePost(post)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err.Error())
		return
	}
	err = db.HandlerDB.ServiceDB.MiddlewareData.InsertPost(post)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, "Internal server error")
		return
	}
}

func (db *HandlerLayer) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
}

func (db *HandlerLayer) GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	post := db.HandlerDB.ServiceDB.MiddlewareData.GetPostById(id)
	fmt.Println(post)
}
