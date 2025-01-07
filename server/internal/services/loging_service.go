package services

import (
	"net/http"
	"time"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (db *ServiceLayer) ValidateLogin(user types.User) (int, bool) {
	if err := utils.ValidateLength(user.Nickname); err != nil {
		return 0, false
	}
	if err := utils.ValidateLength(user.Password); err != nil {
		return 0, false
	}
	id, hash := db.ServiceDB.MiddlewareData.GetUserPassword(user.Nickname)
	return id, utils.ComparePass(user.Password, hash)
}

func (db *ServiceLayer) UpdateSession(w http.ResponseWriter, userId int) error {
	token := utils.GenerateToken()
	currentTime := time.Now()
	err := db.ServiceDB.MiddlewareData.InsertSession(userId, token, currentTime)
	if err != nil {
		return err
	}
	utils.GiveCookie(w, token, currentTime)
	return nil
}
