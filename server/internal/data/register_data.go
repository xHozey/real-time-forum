package data

import "forum/server/internal/types"

func (db *DataLayer) GetUserEmail(email string) bool {
	exists := false
	db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_profile WHERE email = ?)", email).Scan(&exists)
	return exists
}

func (db *DataLayer) GetUserNickname(nickname string) bool {
	exists := false
	db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_profile WHERE nickname = ?)", nickname).Scan(&exists)
	return exists
}

func (db *DataLayer) InsertUser(user types.User) error {
	stm, err := db.DataDB.Prepare("INSERT INTO user_profile (nickname,email,password,gender,firstname,lastname,age) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stm.Exec(user.Nickname, user.Email, user.Password, user.Gender, user.FirstName, user.LastName, user.Age)
	return err
}
