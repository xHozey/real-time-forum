package data

import (
	"time"

	"forum/server/internal/types"
)

func (db *DataLayer) CheckEmailExist(email string) bool {
	exists := false
	db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_profile WHERE email = ?)", email).Scan(&exists)
	return exists
}

func (db *DataLayer) CheckUserExist(nickname string) bool {
	exists := false
	db.DataDB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_profile WHERE nickname = ?)", nickname).Scan(&exists)
	return exists
}

func (db *DataLayer) GetUserPassword(user string) (int, string) {
	var id int
	var password string
	db.DataDB.QueryRow("SELECT id, password FROM user_profile WHERE nickname = ? OR email = ?", user, user).Scan(&id, &password)
	return id, password
}

func (db *DataLayer) InsertUser(user types.User) error {
	stm, err := db.DataDB.Prepare("INSERT INTO user_profile (nickname,email,password,gender,firstname,lastname,age) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stm.Exec(user.Nickname, user.Email, user.Password, user.Gender, user.FirstName, user.LastName, user.Age)
	return err
}

func (db *DataLayer) GetUserBySession(token string) (int, string) {
	var id int
	var nickname string
	db.DataDB.QueryRow("SELECT id, nickname FROM user_profile LEFT JOIN session ON  user_profile.id = session.user_id WHERE session.token = ?", token).Scan(&id, &nickname)
	if !db.ValidateSession(id) {
		return 0, ""
	}
	return id, nickname
}

func (db *DataLayer) InsertSession(id int, token string, now time.Time) error {
	db.DeleteSession(id)
	// reset sesion time to one hour after tests
	_, err := db.DataDB.Exec("INSERT INTO session (user_id,token, expire_at, created_at) VALUES (?,?,?,?)", id, token, now.Add(time.Hour*6), now)
	return err
}

func (db *DataLayer) DeleteSession(id int) {
	db.DataDB.Exec("DELETE FROM session WHERE user_id = ?", id)
}

func (db *DataLayer) ValidateSession(id int) bool {
	var expireTime time.Time
	db.DataDB.QueryRow("SELECT expire_at FROM session LEFT JOIN user_profile ON  session.user_id = user_profile.id WHERE user_profile.id = ?", id).Scan(&expireTime)
	return expireTime.Compare(time.Now()) > 0
}

func (db *DataLayer) GetUserByName(nickname string) int {
	var id int
	db.DataDB.QueryRow("SELECT id FROM user_profile WHERE nickname = ?", nickname).Scan(&id)
	return id
}
