package interfaceUseCase

import (
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
)

type IAdminUseCase interface {
	AdminLogin(*requestmodels.AdminLogin) (*responsemodels.AdminLogin, error)
}
