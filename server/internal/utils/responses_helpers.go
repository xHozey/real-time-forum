package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponseStatus(w http.ResponseWriter, statusCode int, err error) {
	log.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err = json.NewEncoder(w).Encode(err.Error())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func SendJsonData(w http.ResponseWriter, data any) error{
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		SendResponseStatus(w, http.StatusInternalServerError, err)
		return err
	}
	return nil
}
