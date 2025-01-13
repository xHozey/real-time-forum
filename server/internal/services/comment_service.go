package services

import (
	"forum/server/internal/types"
	"strings"
)

func (db *ServiceLayer) ValidateComment(comment types.Comment)error {
	if len(strings.TrimSpace(comment.Content)) == 0 || len(comment.Content) > 2500 {
		return types.ErrInvalidComment
	}
	if !db.ServiceDB.MiddlewareData.CheckIfPostExist(comment.PostId) {
		return types.ErrNotExist
	}
	return nil
}