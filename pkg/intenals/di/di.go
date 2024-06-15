package di

import (
	"fmt"

	server "github.com/Vajid-Hussain/machine-test/pkg/api"
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/db"
)

func InitializeAPI(config *config.Config) (*server.Server, error) {
	_, err := db.ConnectDatabase(&config.DB)
	if err != nil {
		return nil, err
	}

	fmt.Println("hlo")
	return server.InitServer(), nil

}
