package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponseStatus(w http.ResponseWriter, statusCode int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(message)
	return err
}
