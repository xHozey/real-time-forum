package data

import (
	"html"

	"forum/server/internal/types"
)

func (db *DataLayer) InsertComment(comment types.Comment) error {
	_, err := db.DataDB.Exec("INSERT INTO comment (user_id,post_id,content) VALUES (?,?,?)", comment.UserId, comment.PostId, html.EscapeString(comment.Content))
	if err != nil {
		return err
	}
	return nil
}

func (db *DataLayer) GetComments(post_id int) ([]types.Comment, error) {
	comments := []types.Comment{}
	rows, err := db.DataDB.Query(`
    SELECT 
        c.*,
        (SELECT COUNT(*) FROM comment_react WHERE comment_id = c.id AND type = 1) AS likes,
        (SELECT COUNT(*) FROM comment_react WHERE comment_id = c.id AND type = -1) AS dislikes
    FROM comment c WHERE post_id = ?`, post_id)
	if err != nil {
		return []types.Comment{}, err
	}
	for rows.Next() {
		comment := types.Comment{}
		rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreationDate, &comment.Likes, &comment.Dislikes)
		comment.Author = db.GetUserNameById(comment.UserId)
		comment.IsLiked, comment.IsDisliked = db.CheckIfLikedComment(comment.Id, comment.UserId)
		comments = append(comments, comment)
	}
	return comments, nil
}
