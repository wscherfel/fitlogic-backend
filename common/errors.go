package common

import "errors"

// Error messages used in return JSONs
var (
	ErrUserNotFound = errors.New("User with given credentials was not found")

	ErrWrongEmailOrPassword = errors.New("Wrong error and password combination")
)

type Error struct {
	Error string
}

func CreateError(err error) *Error {
	return &Error{Error: err.Error()}
}


