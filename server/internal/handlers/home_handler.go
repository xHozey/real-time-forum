package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func (db *HandlerLayer) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../../client/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
}
