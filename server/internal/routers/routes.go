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
	mux.HandleFunc("/ws", dbLayers.WsHandler)

	return mux
}

func LinkLayers(db *sql.DB) handlers.HandlerLayer {
	dataDb := data.DataLayer{DataDB: db}
	middlewareDb := middleware.MiddleWareLayer{MiddlewareData: dataDb}
	serviceDb := services.ServiceLayer{ServiceDB: middlewareDb}
	handlerDb := handlers.HandlerLayer{HandlerDB: serviceDb}
	return handlerDb
}
