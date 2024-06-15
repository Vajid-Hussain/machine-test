package interfaceRepository

import (
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
)

type IUserRepository interface {
	UserSignUP(*requestmodels.UserSignup) (*responsemodels.UserSignup, error)
	FetchUserPassword(string) (string, error)
	FetchUserID(string) (string, error)
}
