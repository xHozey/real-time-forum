package handlers

import (
	"html/template"
	"log"
	"net/http"

	"forum/server/internal/utils"
)

func (db *HandlerLayer) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../client/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	id, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
	if r.URL.Path == "/" {
		if id == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	} else if r.URL.Path == "/login" || r.URL.Path == "/register" {
		if id != 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	tpl.Execute(w, nil)
}
