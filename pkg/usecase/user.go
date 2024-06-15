package usecase

import (
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceRepository "github.com/Vajid-Hussain/machine-test/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/machine-test/pkg/usecase/interface"
	"github.com/Vajid-Hussain/machine-test/pkg/utils"
)

type userUseCase struct {
	repo  interfaceRepository.IUserRepository
	token config.JWTConfig
}

func NewUserUseCase(repository interfaceRepository.IUserRepository, token config.JWTConfig) interfaceUseCase.IUserUseCase {
	return &userUseCase{repo: repository, token: token}
}

func (r *userUseCase) UserSignUP(userProfile *requestmodels.UserSignup) (*responsemodels.UserSignup, error) {
	// Password hashing
	var err error
	userProfile.Password, err = utils.HashPassword(userProfile.Password)
	if err != nil {
		return nil, err
	}

	// Add User in Database
	res, err := r.repo.UserSignUP(userProfile)
	if err != nil {
		return nil, err
	}

	// Generate autorization token
	res.AccessToken, err = utils.GenerateAcessToken(r.token.SecretKeyUser, res.ID, r.token.ExpirationTime)
	if err != nil {
		return nil, err
	}

	res.RefreshToken, err = utils.GenerateRefreshToken(r.token.SecretKeyUser, res.ID)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (r *userUseCase) UserLogin(loginCredential *requestmodels.UserLogin) (*responsemodels.UserLogin, error) {
	var res responsemodels.UserLogin

	// Get user Password
	passowrd, err := r.repo.FetchUserPassword(loginCredential.Email)
	if err != nil {
		return nil, err
	}

	// compare plain password with encrypted password
	match := utils.CompairPassword(passowrd, loginCredential.Password)
	if match != nil {
		return nil, match
	}

	// Get user id by using email
	userID, err := r.repo.FetchUserID(loginCredential.Email)
	if err != nil {
		return nil, err
	}

	// Generate autorization token
	res.AccessToken, err = utils.GenerateAcessToken(r.token.SecretKeyUser, userID, r.token.ExpirationTime)
	if err != nil {
		return nil, err
	}

	res.RefreshToken, err = utils.GenerateRefreshToken(r.token.SecretKeyUser, userID)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
