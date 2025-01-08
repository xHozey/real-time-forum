package routes

import (
	"database/sql"
	"net/http"
	"time"

	"forum/server/internal/data"
	"forum/server/internal/handlers"
	middleware "forum/server/internal/middleWare"
	"forum/server/internal/services"
)

func Routes(db *sql.DB) *http.ServeMux {
	dbLayers := LinkLayers(db)
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../../client/app"))
	mux.Handle("/app/", http.StripPrefix("/app/", fs))
	mux.HandleFunc("/ws", dbLayers.WsHandler)
	mux.Handle("POST /api/login", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.LoginHandler), 3, time.Second*30))
	mux.Handle("POST /api/register", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.RegisterHandler), 3, time.Second*30))
	mux.Handle("POST /api/addpost", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.PostHandler), 10, time.Second*5))
	mux.Handle("GET /api/post", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.GetPostsHandler), 10, time.Second*30))
	mux.Handle("GET /api/post/{id}", dbLayers.HandlerDB.ServiceDB.RateLimiter(http.HandlerFunc(dbLayers.GetPostByIdHandler), 10, time.Second*30))
	mux.HandleFunc("POST /logout", dbLayers.LogoutHandler)
	return mux
}

func LinkLayers(db *sql.DB) handlers.HandlerLayer {
	dataDb := data.DataLayer{DataDB: db}
	middlewareDb := middleware.MiddleWareLayer{MiddlewareData: dataDb}
	serviceDb := services.ServiceLayer{ServiceDB: middlewareDb}
	handlerDb := handlers.HandlerLayer{HandlerDB: serviceDb}
	return handlerDb
}
