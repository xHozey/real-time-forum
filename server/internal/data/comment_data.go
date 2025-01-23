package data

import (
	"fmt"
	"html"

	"forum/server/internal/types"
)

func (db *DataLayer) InsertComment(comment types.Comment) (types.Comment, error) {
	last, err := db.DataDB.Exec("INSERT INTO comment (user_id,post_id,content) VALUES (?,?,?)", comment.UserId, comment.PostId, html.EscapeString(comment.Content))
	if err != nil {
		return types.Comment{}, err
	}
	lastRow, _ := last.LastInsertId()
	insertedComment := db.getSingleComment(int(lastRow))
	return insertedComment, nil
}

func (db *DataLayer) getSingleComment(comment_id int) types.Comment {
	comment := types.Comment{}
	db.DataDB.QueryRow(`
    SELECT 
        c.*,
        (SELECT COUNT(*) FROM comment_react WHERE comment_id = c.id AND type = 1) AS likes,
        (SELECT COUNT(*) FROM comment_react WHERE comment_id = c.id AND type = -1) AS dislikes
    FROM comment c WHERE id = ?`, comment_id).Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreationDate, &comment.Likes, &comment.Dislikes)
	return comment
}

func (db *DataLayer) GetComments(post_id, offset int) ([]types.Comment, error) {
	comments := []types.Comment{}
	rows, err := db.DataDB.Query(`
    SELECT 
        c.*,
        (SELECT COUNT(*) FROM comment_react WHERE comment_id = c.id AND type = 1) AS likes,
        (SELECT COUNT(*) FROM comment_react WHERE comment_id = c.id AND type = -1) AS dislikes
    FROM comment c WHERE post_id = ? ORDER BY c.created_at DESC LIMIT ? OFFSET ?`, post_id, types.Limit, offset)
	if err != nil {
		fmt.Println(err)
		return []types.Comment{}, err
	}
	for rows.Next() {
		comment := types.Comment{}
		rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreationDate, &comment.Likes, &comment.Dislikes)
		comment.Author = db.GetUserNameById(comment.UserId)
		comment.IsLiked = db.CheckIfLikedComment(comment.Id, comment.UserId)
		comments = append(comments, comment)
	}
	return comments, nil
}
