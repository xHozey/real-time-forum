package data

import (
	"html"

	"forum/server/internal/types"
)

func (db *DataLayer) InsertPost(postData types.Post) error {
	res, err := db.DataDB.Exec("INSERT INTO post (user_id,content) VALUES (?,?)", postData.UserId, html.EscapeString(postData.Content))
	if err != nil {
		return err
	}
	post_id, err := res.LastInsertId()
	for _, val := range postData.Categories {
		id := db.GetCategorieId(val)
		db.InsertPostCategorie(post_id, id)
	}
	return err
}

func (db *DataLayer) CheckIfPostExist(post_id int) bool {
	exists := false
	db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM post WHERE id = ?)", post_id).Scan(&exists)
	return exists
}

func (db *DataLayer) InsertPostCategorie(post_id int64, categorie_id int) error {
	_, err := db.DataDB.Exec("INSERT INTO post_category (post_id,category_id) VALUES (?,?)", post_id, categorie_id)
	return err
}

func (db *DataLayer) GetCategorieId(categorie string) int {
	var id int
	db.DataDB.QueryRow("SELECT id FROM category WHERE category_name = ?", categorie).Scan(&id)
	return id
}

func (db *DataLayer) GetAllPosts() ([]types.Post, error) {
	posts := []types.Post{}
	rows, err := db.DataDB.Query(`
    SELECT 
        p.*,
        (SELECT COUNT(*) FROM post_react WHERE post_id = p.id AND type = 1) AS likes,
        (SELECT COUNT(*) FROM post_react WHERE post_id = p.id AND type = -1) AS dislikes
    FROM post p`)
	if err != nil {
		return []types.Post{}, err
	}
	for rows.Next() {
		post := types.Post{}
		rows.Scan(&post.Id, &post.UserId, &post.Content, &post.CreationDate, &post.Likes, &post.Dislikes)
		rows, err := db.DataDB.Query("SELECT category_name FROM category c LEFT JOIN post_category p ON c.id = p.category_id WHERE p.post_id = ?", post.Id)
		if err != nil {
			return []types.Post{}, err
		}
		for rows.Next() {
			category := ""
			rows.Scan(&category)
			post.Categories = append(post.Categories, category)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (db *DataLayer) GetPostById(id int) (types.Post, error) {
	post := types.Post{}
	err := db.DataDB.QueryRow(`
    SELECT 
        p.*,
        (SELECT COUNT(*) FROM post_react WHERE post_id = p.id AND type = 1) AS likes,
        (SELECT COUNT(*) FROM post_react WHERE post_id = p.id AND type = -1) AS dislikes
    FROM post p
    WHERE p.id = ?`, id).Scan(&post.Id, &post.UserId, &post.Content, &post.CreationDate, &post.Likes, &post.Dislikes)
	if err != nil {
		return types.Post{}, err
	}
	rows, err := db.DataDB.Query("SELECT category_name FROM category c LEFT JOIN post_category p ON c.id = p.category_id WHERE p.post_id = ?", post.Id)
	if err != nil {
		return types.Post{}, err
	}
	for rows.Next() {
		category := ""
		rows.Scan(&category)
		post.Categories = append(post.Categories, category)
	}
	return post, nil
}
