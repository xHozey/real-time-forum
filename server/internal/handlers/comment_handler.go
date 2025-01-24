package handlers

import (
	"log"
	"net/http"
	"strconv"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
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
	comments, err := db.HandlerDB.ServiceDB.MiddlewareData.GetComments(id, offset)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
	err = utils.SendJsonData(w, &comments)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}

func (db *HandlerLayer) CommentHandler(w http.ResponseWriter, r *http.Request) {
	comment := types.Comment{}
	err := utils.DecodeRequest(r, &comment)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusBadRequest, types.ErrBadRequest)
		return
	}
	id, nickname := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	comment.UserId = id
	if comment.UserId == 0 {
		utils.SendResponseStatus(w, http.StatusUnauthorized, types.ErrUnauthorized)
		return
	}
	if err := db.HandlerDB.ValidateComment(comment); err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	insertedComment, err := db.HandlerDB.ServiceDB.MiddlewareData.InsertComment(comment)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
	insertedComment.Author = nickname
	err = utils.SendJsonData(w, &insertedComment)
	if err != nil {
		log.Println(err)
		utils.SendResponseStatus(w, http.StatusInternalServerError, types.ErrInternalServerError)
		return
	}
}
