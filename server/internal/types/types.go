package types

import (
	"time"
)

const (
	CookieName = "session_cookie"
	Limit      = 10
)

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
	Author       string    `json:"author"`
	UserId       int       `json:"user_id"`
	Content      string    `json:"content"`
	Categories   []string  `json:"categories"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	CreationDate time.Time `json:"creation_date"`
	IsLiked      int       `json:"isliked"`
}

type Comment struct {
	Id           int       `json:"id"`
	Author       string    `json:"author"`
	PostId       int       `json:"post_id"`
	UserId       int       `json:"user_id"`
	Content      string    `json:"content"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	CreationDate time.Time `json:"creation_date"`
	IsLiked      int       `json:"isliked"`
}

type Reaction struct {
	Thread_type string `json:"thread_type"`
	Thread_id   int    `json:"thread_id"`
	Reaction    int    `json:"reaction"`
}

type Clients struct {
	UserId   int    `json:"id"`
	Nickname string `json:"nickname"`
	Status   bool   `json:"status"`
}

type InfoUser struct {
	UserId   int
	Nickname string
	Clients  []Clients
}

type Messages struct {
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Content  string    `json:"content"`
	Creation time.Time `json:"creation"`
}

type Scroll struct {
	Limit  int `json:"limit"`
	OffSet int `json:"offset"`
}
