package routes

import (
	"forum/server/internal/handlers"
	"net/http"
)

func wsRoute(mux *http.ServeMux) {
	mux.HandleFunc("/ws", handlers.WsHandler)
}