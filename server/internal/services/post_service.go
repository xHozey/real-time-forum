package services

import (
	"strings"

	"forum/server/internal/types"
)

func (db *ServiceLayer) ValidatePost(post types.Post) error {
	if len(strings.TrimSpace(post.Content)) == 0 || len(post.Content) > 500 {
		return types.ErrInvalidPost
	}
	if !checkDuplicate(post.Categories) {
		return types.ErrInvalidCategorie
	}
	for _, val := range post.Categories {
		id := db.ServiceDB.MiddlewareData.GetCategorieId(val)
		if id == 0 {
			return types.ErrInvalidCategorie
		}
	}
	return nil
}

func checkDuplicate(data []string) bool {
	hashMap := make(map[string]bool)
	for _, val := range data {
		if !hashMap[val] {
			hashMap[val] = true
		} else {
			return false
		}
	}
	return true
}
