package common

import "errors"

// Error messages used in return JSONs
var (
	ErrWrongEmailOrPassword = errors.New("Wrong email and password combination")

	ErrMissingTokenClaims = errors.New("Sent token is valid but has missing claims. Try to log in again to obtain a new one")

	ErrUnsufficientPrivileges = errors.New("Logged user does not have sufficient privileges to do this operation.")

	ErrIdInPathWrongFormat = errors.New("ID in path is not a valid ID")

	ErrCannotCreateProjectForOthers = errors.New("Cannot create project with other user as project manager")

	ErrWrongPassword = errors.New("Wrong password")

	ErrManagerStillLeadsProjects = errors.New("Manager you want to downgrade still leads a project")

	ErrDateOutOfRange = errors.New("Input date out of range")

	ErrStartDateAfterEnd = errors.New("Start date is the same or after end date")

	ErrCannotDeleteOnlyAdmin = errors.New("Cannot delete only administrator")
)

// Error is a structure of error message returned in json
type Error struct {
	Error string
}

// CreateError will create error structure for return JSON
// from golang error
func CreateError(err error) *Error {
	return &Error{Error: err.Error()}
}


