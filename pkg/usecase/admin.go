package usecase

import (
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceRepository "github.com/Vajid-Hussain/machine-test/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/machine-test/pkg/usecase/interface"
	"github.com/Vajid-Hussain/machine-test/pkg/utils"
)

type adminUseCase struct {
	repo  interfaceRepository.IAdminRepository
	token config.JWTConfig
}

func NewAdminUseCase(repository interfaceRepository.IAdminRepository, token config.JWTConfig) interfaceUseCase.IAdminUseCase {
	return &adminUseCase{repo: repository, token: token}
}

func (r *adminUseCase) AdminLogin(loginCredential *requestmodels.AdminLogin) (*responsemodels.AdminLogin, error) {
	var res responsemodels.AdminLogin

	// Get Admin Password
	passowrd, err := r.repo.FetchAdminPassword(loginCredential.Email)
	if err != nil {
		return nil, err
	}

	// compare plain password with encrypted password
	match := utils.CompairPassword(passowrd, loginCredential.Password)
	if match != nil {
		return nil, match
	}

	// Get Admin id by using email
	adminID, err := r.repo.FetchAdminID(loginCredential.Email)
	if err != nil {
		return nil, err
	}

	// Generate autorization token
	res.AccessToken, err = utils.GenerateAcessToken(r.token.SecretKeyAdmin, adminID, r.token.ExpirationTime)
	if err != nil {
		return nil, err
	}

	res.RefreshToken, err = utils.GenerateRefreshToken(r.token.SecretKeyAdmin, adminID)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
