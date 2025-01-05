package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest(r *http.Request, receiver any) error {
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(&receiver)
	if err != nil {
		return err
	}
	return nil
}
