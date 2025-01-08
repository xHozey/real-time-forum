package data

import "forum/server/internal/types"

func (db *DataLayer) InsertPost(postData types.Post) error {
	res, err := db.DataDB.Exec("INSERT INTO post (user_id,content) VALUES (?,?)", postData.User, postData.Content)
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

func (db *DataLayer) InsertPostCategorie(post_id int64, categorie_id int) error {
	_, err := db.DataDB.Exec("INSERT INTO post_category (post_id,category_id) VALUES (?,?)", post_id, categorie_id)
	return err
}

func (db *DataLayer) GetCategorieId(categorie string) int {
	var id int
	db.DataDB.QueryRow("SELECT id FROM category WHERE category_name = ?", categorie).Scan(&id)
	return id
}
