package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"forum/server/internal/types"
)

func DecodeRequest(r *http.Request, receiver any) error {
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(&receiver)
	if err != nil {
		return err
	}
	return nil
}

func GetCookie(r *http.Request) string {
	cookie, err := r.Cookie(types.CookieName)
	if err == http.ErrNoCookie {
		return ""
	}
	return cookie.Value
}

func GiveCookie(w http.ResponseWriter, token string, now time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     types.CookieName,
		Value:    token,
		Expires:  now.Add(time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
}

func DeleteCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		MaxAge: -1,
	})
}
