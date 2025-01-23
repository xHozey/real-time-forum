package handlers

import (
	"net/http"
	"strconv"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendResponseStatus(w, http.StatusNotFound, err)
		return
	}
	offset, err := utils.GetOffset(w,r)
	if err != nil {
		return
	}
	comments, err := db.HandlerDB.ServiceDB.MiddlewareData.GetComments(id, offset)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
		return
	}
	err = utils.SendJsonData(w, &comments)
	if err != nil {
		return
	}
}

func (db *HandlerLayer) CommentHandler(w http.ResponseWriter, r *http.Request) {
	comment := types.Comment{}
	err := utils.DecodeRequest(r, &comment)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	comment.UserId, _ = db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	if comment.UserId == 0 {
		utils.SendResponseStatus(w, http.StatusUnauthorized, types.ErrUnauthorized)
		return
	}
	if err := db.HandlerDB.ValidateComment(comment); err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	insertedComment, err := db.HandlerDB.ServiceDB.MiddlewareData.InsertComment(comment)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusLocked, err)
		return
	}
	utils.SendJsonData(w, &insertedComment)
}
