package responsemodels

import "errors"

// User Repository
var (
	ErrInternalServer = errors.New("an unexpected error occurred. Please try again later")
	ErrEmailExist     = errors.New("email provided is already associated with an existing account. Could you please try using a different email?")
	ErrNoEmailExist   = errors.New("no account found with the provided email address. Please use a different email")
)

// User Usecase
var (
	ErrWrongContecntTypeResume = errors.New("wrong content type")
)

// Job Repository
var (
	ErrNoActiveJob       = errors.New("no active job's are available")
	ErrJobAlreadyApplied = errors.New("job alredy applied or no job exist")
	ErrNoAppliedJob = errors.New("you don't have any applied job")
)

// Admin usecase
var (
	ErrPaginationWrongValue = errors.New("pagination value must be posiive")
)
