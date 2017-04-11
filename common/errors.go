package common

import "errors"

// Error messages used in return JSONs
var (
	ErrWrongEmailOrPassword = errors.New("Wrong error and password combination")

	ErrMissingTokenClaims = errors.New("Sent token is valid but has missing claims. Try to log in again to obtain a new one")

	ErrUnsufficientPrivileges = errors.New("Logged user does not have sufficient privileges to do this operation.")

	ErrIdInPathWrongFormat = errors.New("ID in path is not a valid ID")
)

type Error struct {
	Error string
}

func CreateError(err error) *Error {
	return &Error{Error: err.Error()}
}


