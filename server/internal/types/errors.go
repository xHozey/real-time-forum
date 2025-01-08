package types

import "errors"

var (
	ErrEmailAlreadyTaken    = errors.New("this email already taken.Please choose another email")
	ErrNicknameAlreadyTaken = errors.New("this nickname already taken.Please choose another nickname")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidNickname      = errors.New("invalid nickname")
	ErrInvalidFirstName     = errors.New("invalid first name")
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrIncorrectGender      = errors.New("please choose male or female")
	ErrIncorrectLength      = errors.New("characters must be between 3 and 32")
	ErrAgeLimit             = errors.New("incorrect age")
	ErrInvalidCategorie     = errors.New("invalid categorie")
	ErrInvalidPost          = errors.New("invalid post")
)

const CookieName = "session_cookie"
