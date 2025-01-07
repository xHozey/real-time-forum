package utils

import (
	"forum/server/internal/types"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePass(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func ValidateLength(data string) error {
	if len(data) <= 3 || len(data) >= 32 {
		return types.ErrIncorrectLength
	}
	return nil
}
