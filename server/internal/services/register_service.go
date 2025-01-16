package services

import (
	"net/mail"
	"strings"

	"forum/server/internal/types"
	"forum/server/internal/utils"
)

func (service *ServiceLayer) ValidateCredentials(user types.User) error {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return err
	}
	if err := validateGender(user.Gender); err != nil {
		return err
	}
	if err := utils.ValidateLength(user.FirstName); err != nil {
		return err
	}
	if err := utils.ValidateLength(user.LastName); err != nil {
		return err
	}
	if err := utils.ValidateLength(user.Nickname); err != nil {
		return err
	}
	if err := utils.ValidateLength(user.Password); err != nil {
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
	if exists := db.ServiceDB.MiddlewareData.CheckEmailExist(strings.ToLower(user.Email)); exists {
		return types.ErrEmailAlreadyTaken
	}
	if exists := db.ServiceDB.MiddlewareData.CheckUserExist(strings.ToLower(user.Nickname)); exists {
		return types.ErrNicknameAlreadyTaken
	}
	return nil
}

func (db *ServiceLayer) InsertClearCredential(user types.User) error {
	pass, err := utils.HashPass(user.Password)
	if err != nil {
		return err
	}
	user.Password = pass
	user.Nickname = strings.ToLower(user.Nickname)
	user.Email = strings.ToLower(user.Email)
	err = db.ServiceDB.MiddlewareData.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func validateGender(gender string) error {
	if gender != "male" && gender != "female" {
		return types.ErrIncorrectGender
	}
	return nil
}
