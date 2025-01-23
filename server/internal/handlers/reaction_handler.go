package handlers

import (
	"net/http"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *HandlerLayer) ReactionHandler(w http.ResponseWriter, r *http.Request) {
	react := types.Reaction{}
	err := utils.DecodeRequest(r, &react)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}
	id, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))

	err = db.HandlerDB.CheckReactInput(react)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusBadRequest, err)
		return
	}

	err = db.HandlerDB.ReactionService(react, id)
	if err != nil {
		utils.SendResponseStatus(w, http.StatusInternalServerError, err)
		return
	}
}
