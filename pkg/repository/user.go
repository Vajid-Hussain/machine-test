package repository

import (
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceRepository "github.com/Vajid-Hussain/machine-test/pkg/repository/interface"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaceRepository.IUserRepository {
	return &userRepository{DB: db}
}

func (d *userRepository) UserSignUP(userProfile *requestmodels.UserSignup) (*responsemodels.UserSignup, error) {
	var response responsemodels.UserSignup

	query := `INSERT INTO users (name, email, password, address, user_type, profile_headline ,user_since) SELECT $1, $2, $3 , $4, 'User', $5, now()
	WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = $2) RETURNING id, name, email, address, user_type, profile_headline, user_since`
	result := d.DB.Raw(query, userProfile.Name, userProfile.Email, userProfile.Password, userProfile.Address, userProfile.ProfileHeadline).Scan(&response)
	if result.Error != nil {
		return nil, responsemodels.ErrInternalServer
	}
	if result.RowsAffected == 0 {
		return nil, responsemodels.ErrEmailExist
	}

	return &response, nil
}

func (d *userRepository) FetchUserPassword(email string) (password string, err error) {

	query := "SELECT password FROM users WHERE email = $1 AND status='active'"
	result := d.DB.Raw(query, email).Scan(&password)
	if result.Error != nil {
		return "", responsemodels.ErrInternalServer
	}
	if result.RowsAffected == 0 {
		return "", responsemodels.ErrNoEmailExist
	}

	return password, nil
}

func (d *userRepository) FetchUserID(email string) (id string, err error) {

	query := "SELECT id FROM users WHERE email = $1 AND status='active'"
	result := d.DB.Raw(query, email).Scan(&id)
	if result.Error != nil {
		return "", responsemodels.ErrInternalServer
	}
	if result.RowsAffected == 0 {
		return "", responsemodels.ErrNoEmailExist
	}

	return id, nil
}
