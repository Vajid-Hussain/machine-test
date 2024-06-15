package main

import (
	"log"

	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/di"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal("config err ",err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("error initialize api ",err)
	}

	server.Start(config.Server.Port)
}
