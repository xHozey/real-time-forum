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
	if r.URL.Path == "/" {
		id, _ := db.HandlerDB.ServiceDB.MiddlewareData.GetUserBySession(utils.GetCookie(r))
		if id == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
	tpl.Execute(w, nil)
}
