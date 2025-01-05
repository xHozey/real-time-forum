package services

import (
	"html"
	"net/mail"
	"forum/server/internal/types"
)

func (service *ServiceLayer) ValidateCredentials(user types.User) error {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return err
	}
	if err := validateGender(user.Gender); err != nil {
		return err
	}
	if err := validateLength(user.FirstName); err != nil {
		return err
	}
	if err := validateLength(user.LastName); err != nil {
		return err
	}
	if err := validateLength(user.Nickname); err != nil {
		return err
	}
	if err := validateLength(user.Password); err != nil {
		return err
	}
	if user.Age > 100 || user.Age <= 0 {
		return types.ErrAgeLimit
	}
	if err := service.checkIfExists(user); err != nil {
		return err
	}
	return nil
}

func (db *ServiceLayer) checkIfExists(user types.User) error {
	if exists := db.ServiceDB.GetUserEmail(user.Email); exists {
		return types.ErrEmailAlreadyTaken
	}
	if exists := db.ServiceDB.GetUserNickname(user.Nickname); exists {
		return types.ErrNicknameAlreadyTaken
	}
	return nil
}

func (db *ServiceLayer) InsertClearCredential(user types.User) {
	user.Nickname = html.EscapeString(user.Nickname)
	user.LastName = html.EscapeString(user.LastName)
	user.FirstName = html.EscapeString(user.FirstName)
	user.Email = html.EscapeString(user.Email)
	db.ServiceDB.InsertUser(user)
}

func validateGender(gender string) error {
	if gender != "male" && gender != "female" {
		return types.ErrIncorrectGender
	}
	return nil
}

func validateLength(data string) error {
	if len(data) <= 3 || len(data) >= 32 {
		return types.ErrIncorrectLength
	}
	return nil
}