package routes

import (
	"database/sql"
	"net/http"

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
	mux.Handle("/", middleware.MethodMiddleware(http.HandlerFunc(dbLayers.HomeHandler), http.MethodGet))
	mux.Handle("/login", middleware.MethodMiddleware(http.HandlerFunc(dbLayers.LoginHandler), http.MethodPost))
	mux.Handle("/register", middleware.MethodMiddleware(http.HandlerFunc(dbLayers.RegisterHandler), http.MethodPost))
	mux.Handle("/logout", middleware.MethodMiddleware(http.HandlerFunc(dbLayers.LogoutHandler), http.MethodPost))
	mux.HandleFunc("/ws", dbLayers.WsHandler)
	return mux
}

func LinkLayers(db *sql.DB) handlers.HandlerLayer {
	dataDb := data.DataLayer{DataDB: db}
	serviceDb := services.ServiceLayer{ServiceDB: dataDb}
	handlerDb := handlers.HandlerLayer{HandlerDB: serviceDb}
	return handlerDb
}
