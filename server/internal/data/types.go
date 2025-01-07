package data

import "database/sql"

type DataLayer struct {
	DataDB *sql.DB
}

type DataMethods interface {
	GetUserEmail()
	GetUserNickname()
	GetUserPassword()
	InsertUser()
	GetUserBySession()
	InsertSession()
	DeleteSession()
	ValidateSession()
	GiveBucket()
	RefillTokens()
	ExtractBucketDate()
	TakeToken()
}
