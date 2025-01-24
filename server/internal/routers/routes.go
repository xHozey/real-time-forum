package routes

import (
	"database/sql"
	"net/http"
	"time"

	"forum/server/internal/data"
	"forum/server/internal/handlers"
	middleware "forum/server/internal/middleWare"
	"forum/server/internal/services"
	"forum/server/internal/websocket"
)

func Routes(db *sql.DB) *http.ServeMux {
	dbLayers, wsLayer := LinkLayers(db)
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../client/app"))
	mux.Handle("/app/", http.StripPrefix("/app/", fs))
	mux.HandleFunc("/", dbLayers.HomeHandler)
	mux.HandleFunc("/ws", wsLayer.WsHandler)
	mux.Handle("GET /api/client/{id}", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.ClientHandler), 100, time.Second*5))
	mux.Handle("POST /api/login", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.LoginHandler), 3, time.Second*30))
	mux.Handle("POST /api/register", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.RegisterHandler), 3, time.Second*30))
	mux.Handle("POST /api/addpost", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.PostHandler), 10, time.Second*10))
	mux.Handle("POST /api/addcomment", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.CommentHandler), 10, time.Second*10))
	mux.Handle("GET /api/post", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.GetPostsHandler), 100, time.Second*10))
	mux.Handle("GET /api/post/{id}", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.GetPostByIdHandler), 100, time.Second*10))
	mux.Handle("GET /api/post/{id}/comment", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.GetCommentHandler), 100, time.Second*10))
	mux.Handle("POST /api/reaction", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.ReactionHandler), 40, time.Second))
	mux.HandleFunc("/api/info", dbLayers.InfoHandler)
	mux.HandleFunc("GET /logout", dbLayers.LogoutHandler)
	return mux
}

func LinkLayers(db *sql.DB) (handlers.HandlerLayer, websocket.WSlayer) {
	dataDb := data.DataLayer{DataDB: db}
	middlewareDb := middleware.MiddleWareLayer{MiddlewareData: dataDb}
	serviceDb := services.ServiceLayer{ServiceDB: middlewareDb}
	handlerDb := handlers.HandlerLayer{HandlerDB: serviceDb}
	wsData := websocket.WSlayer{Data: dataDb}
	return handlerDb, wsData
}
