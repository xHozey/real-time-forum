package main

import (
	"log"
	"net/http"

	data "forum/server/internal/data"
	routes "forum/server/internal/routers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := data.InitDb()
	if err != nil {
		log.Println(err)
	}

	config := http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(db),
	}
	log.Fatal(config.ListenAndServe())
}
