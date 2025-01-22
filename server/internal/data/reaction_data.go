package data

import (
	"forum/server/internal/types"
)

func (db *DataLayer) CheckPostReaction(user_id, post_id int) (bool, error) {
	var exists bool
	err := db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM post_react WHERE user_id = ? AND post_id = ?)", user_id, post_id).Scan(&exists)
	return exists, err
}

func (db *DataLayer) CheckCommentReaction(user_id, comment_id int) (bool, error) {
	var exists bool
	err := db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM comment_react WHERE user_id = ? AND comment_id = ?)", user_id, comment_id).Scan(&exists)
	return exists, err
}

func (db *DataLayer) DeleteReactionPost(user_id, post_id int) error {
	_, err := db.DataDB.Exec("DELETE FROM post_react WHERE post_id = ? AND user_id = ?", post_id, user_id)
	return err
}

func (db *DataLayer) DeleteReactionComment(user_id, post_id int) error {
	_, err := db.DataDB.Exec("DELETE FROM comment_react WHERE comment_id = ? AND user_id = ?", post_id, user_id)
	return err
}

func (db *DataLayer) GetReactionTypePost(user_id, post_id int) (int, error) {
	var isLiked int
	err := db.DataDB.QueryRow("SELECT type FROM post_react WHERE user_id = ? AND post_id = ?", user_id, post_id).Scan(&isLiked)
	return isLiked, err
}

func (db *DataLayer) GetReactionTypeComment(user_id, post_id int) (int, error) {
	var isLiked int
	err := db.DataDB.QueryRow("SELECT type FROM comment_react WHERE user_id = ? AND comment_id = ?", user_id, post_id).Scan(&isLiked)
	return isLiked, err
}

func (db *DataLayer) InsertReactPost(user_id, post_id, like_type int) error {
	_, err := db.DataDB.Exec("INSERT INTO post_react (post_id, user_id, type) VALUES (?,?,?)", post_id, user_id, like_type)
	return err
}

func (db *DataLayer) InsertReactComment(user_id, post_id, like_type int) error {
	_, err := db.DataDB.Exec("INSERT INTO comment_react (comment_id, user_id, type) VALUES (?,?,?)", post_id, user_id, like_type)
	return err
}

func (db *DataLayer) CheckIfLikedPost(post_id, user_id int) int {
	isLiked := 0
	db.DataDB.QueryRow("SELECT type FROM post_react WHERE post_id = ? AND user_id = ?", post_id, user_id).Scan(&isLiked)
	return isLiked
}

func (db *DataLayer) CheckIfLikedComment(post_id, user_id int) int {
	isLiked := 0
	db.DataDB.QueryRow("SELECT type FROM comment_react WHERE comment_id = ? AND user_id = ?", post_id, user_id).Scan(&isLiked)
	return isLiked
}

func (db *DataLayer) CheckIfThreadExists(react types.Reaction) bool {
	exist := false
	if react.Thread_type == "post" {
		db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM post WHERE id = ?)", react.Thread_id).Scan(&exist)
		return exist
	} else {
		db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM comment WHERE id = ?)", react.Thread_id).Scan(&exist)
		return exist
	}
}
