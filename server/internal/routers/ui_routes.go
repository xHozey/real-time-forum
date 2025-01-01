package routes

import (
	"net/http"

	"forum/server/internal/handlers"
)

func uiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", handlers.HomeHandler)
}
