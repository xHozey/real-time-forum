package main

import (
	"net/http"

	routes "forum/server/internal/routers"
)

func main() {
	config := http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(),
	}
	config.ListenAndServe()
}
