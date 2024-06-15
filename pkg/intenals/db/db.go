package db

import (
	"database/sql"
	"fmt"

	"github.com/Vajid-Hussain/machine-test/pkg/domain"
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	"github.com/Vajid-Hussain/machine-test/pkg/utils"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(DB *config.SqlDatabase) (*gorm.DB, error) {
	// Connecting to postgres sql
	connectionString := fmt.Sprintf("host= %s user= %s  password= %s port= %s sslmode=disable", DB.Host, DB.User, DB.Password, DB.Port)
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Checking database is exist else create new one
	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + DB.Name + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows != nil && rows.Next() {
		fmt.Println("Database " + DB.Name + " already exists.")
	} else {
		_, err := sql.Exec("CREATE DATABASE " + DB.Name)
		if err != nil {
			return nil, err
		}
	}

	// Postgres connection through gorm
	gormConnection, err := gorm.Open(postgres.Open(DB.URL))
	if err != nil {
		return nil, err
	}

	// Creating table by using automigrate
	err = gormConnection.AutoMigrate(
		domain.Users{},
		domain.UserResumeData{},
		domain.Jobs{},
		domain.JobApply{},
	)
	if err != nil {
		return nil, err
	}

	// Explicitly create a initial admin
	err = CheckAndCreateAdmin(gormConnection)
	if err != nil {
		return nil, err
	}

	return gormConnection, nil
}

func CheckAndCreateAdmin(DB *gorm.DB) error {
	var (
		Name     = "sSynlabs"
		Email    = "synlabs@gmail.com"
		Password = "synlabs"
	)

	// Conver plain password to encrypted password
	HashedPassword, err := utils.HashPassword(Password)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (name, email, password, user_type) SELECT $1, $2 , $3 ,'Admin'
	WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = $2) `
	result := DB.Exec(query, Name, Email, HashedPassword)
	if result.Error != nil {
		return responsemodels.ErrInternalServer
	}

	return nil
}
