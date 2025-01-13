package data

func (db *DataLayer) InsertUserMessages(sender int, receiver int, message string) error {
	_, err := db.DataDB.Exec("INSERT INTO message (sender,receiver,content) VALUES (?,?,?)", sender, receiver, message)
	return err
}

func (db *DataLayer) GetMessages(sender int, receiver int, offset int) {
	db.DataDB.Query("SELECT * FROM message WHERE sender = ? AND receiver = ? ")
}
