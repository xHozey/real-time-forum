package routes

import (
	"net/http"

	"forum/server/internal/handlers"
)

func appRoutes(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("../../client/app"))
	mux.Handle("/app/", http.StripPrefix("/app/", fs)) 
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/ws", handlers.WsHandler)
}
