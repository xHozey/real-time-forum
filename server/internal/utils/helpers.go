package utils

import (
	"net/http"
	"strconv"

	"forum/server/internal/types"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePass(pass string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err
}

func ValidateLength(data string) error {
	if len(data) <= 3 || len(data) >= 32 {
		return types.ErrIncorrectLength
	}
	return nil
}

func GenerateToken() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

func GetOffset(w http.ResponseWriter, r *http.Request) (int, error) {
	query := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(query)
	if err != nil {
		SendResponseStatus(w, http.StatusNotFound, err)
		return 0, err
	}
	return offset, nil
}
