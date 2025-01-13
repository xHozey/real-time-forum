package types

import (
	"time"
)

const CookieName = "session_cookie"

type User struct {
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Post struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	Content      string    `json:"content"`
	Categories   []string  `json:"categories"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	CreationDate time.Time `json:"creation_date"`
}

type Comment struct {
	Id           int       `json:"id"`
	PostId       int       `json:"post_id"`
	UserId       int       `json:"user_id"`
	Content      string    `json:"content"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	CreationDate time.Time `json:"creation_date"`
}

type Reaction struct {
	Thread_type string `json:"thread_type"`
	Thread_id   int    `json:"thread_id"`
	Reaction    int    `json:"reaction"`
}

type InfoUser struct {
	User_id  int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Status   bool   `json:"status"`
}

type Message struct {
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Message string `json:"message"`
}