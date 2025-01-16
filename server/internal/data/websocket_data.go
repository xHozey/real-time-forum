package data

import "forum/server/internal/types"

func (db *DataLayer) InsertUserMessages(sender int, receiver int, message string) error {
	_, err := db.DataDB.Exec("INSERT INTO message (sender,receiver,content) VALUES (?,?,?)", sender, receiver, message)
	return err
}

func (db *DataLayer) GetConverceation(source, target int) ([]types.Messages, error) {
	rows, err := db.DataDB.Query(`SELECT sender, receiver, content, created_at 
	          FROM message 
	          WHERE (sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?) 
	          ORDER BY created_at ASC`, source, target, target, source)
	if err != nil {
		return nil, err
	}
	messages := []types.Messages{}
	for rows.Next() {
		message := types.Messages{}
		var sender int
		var receiver int
		rows.Scan(&sender, &receiver, &message.Content, &message.Creation)
		message.Sender = db.GetUserNameById(sender)
		message.Receiver = db.GetUserNameById(receiver)
		messages = append(messages, message)
	}
	return messages, nil
}
