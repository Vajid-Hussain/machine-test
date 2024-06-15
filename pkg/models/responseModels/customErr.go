package responsemodels

import "errors"

// User Repository
var (
	ErrInternalServer = errors.New("an unexpected error occurred. Please try again later")
	ErrEmailExist     = errors.New("email provided is already associated with an existing account. Could you please try using a different email?")
	ErrNoEmailExist   = errors.New("no account found with the provided email address. Please use a different email")
)

// Inventory Repository
var ()

// Admin usecase
var (
	ErrPaginationWrongValue = errors.New("pagination value must be posiive")
)
