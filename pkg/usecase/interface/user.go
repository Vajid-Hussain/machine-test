package interfaceUseCase

import (
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
)

type IUserUseCase interface {
	UserSignUP(*requestmodels.UserSignup) (*responsemodels.UserSignup, error)
	UserLogin(*requestmodels.UserLogin) (*responsemodels.UserLogin, error)
}
