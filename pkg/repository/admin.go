package repository

import (
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceRepository "github.com/Vajid-Hussain/machine-test/pkg/repository/interface"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaceRepository.IAdminRepository {
	return &AdminRepository{DB: db}
}

func (d *AdminRepository) FetchAdminPassword(email string) (password string, err error) {

	query := "SELECT password FROM users WHERE email = $1 AND user_type = 'Admin'"
	result := d.DB.Raw(query, email).Scan(&password)
	if result.Error != nil {
		return "", responsemodels.ErrInternalServer
	}
	if result.RowsAffected == 0 {
		return "", responsemodels.ErrNoEmailExist
	}

	return password, nil
}

func (d *AdminRepository) FetchAdminID(email string) (id string, err error) {

	query := "SELECT id FROM users WHERE email = $1 AND user_type = 'Admin'"
	result := d.DB.Raw(query, email).Scan(&id)
	if result.Error != nil {
		return "", responsemodels.ErrInternalServer
	}
	if result.RowsAffected == 0 {
		return "", responsemodels.ErrNoEmailExist
	}

	return id, nil
}
