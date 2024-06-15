package di

import (
	"fmt"

	server "github.com/Vajid-Hussain/machine-test/pkg/api"
	"github.com/Vajid-Hussain/machine-test/pkg/api/handler"
	"github.com/Vajid-Hussain/machine-test/pkg/api/middlewire"
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/db"
	"github.com/Vajid-Hussain/machine-test/pkg/repository"
	"github.com/Vajid-Hussain/machine-test/pkg/usecase"
)

func InitializeAPI(config *config.Config) (*server.Server, error) {
	DB, err := db.ConnectDatabase(&config.DB)
	if err != nil {
		return nil, err
	}

	// Dipendency Injection
	adminRepository := repository.NewAdminRepository(DB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, config.JWT)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, config.JWT)
	userHandler := handler.NewUserHandler(userUseCase)

	jobRepository := repository.NewJobRepository(DB)
	jobUseCase := usecase.NewJobUseCase(jobRepository, config.S3)
	jobHandler := handler.NewJobHandler(jobUseCase)

	middlewire.NewMiddewire(config.JWT)

	fmt.Println("hlo")
	return server.InitServer(adminHandler, userHandler, jobHandler), nil

}
