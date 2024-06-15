package main

import (
	"log"

	_ "github.com/Vajid-Hussain/machine-test/docs"
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	"github.com/Vajid-Hussain/machine-test/pkg/intenals/di"
)

// @title          Machine task
// @version        1.0
// @description    This is a sample fiber project server.
// @termsOfService http://swagger.io/terms/

// @BasePath /

// @securityDefinitions.apikey authorization
// @in header
// @name authorization
// @description IMPORTANT: TYPE "Bearer" FOLLOWED BY A SPACE AND JWT TOKEN.

func main() {
	
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal("config err ", err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("error initialize api ", err)
	}

	server.Start(config.Server.Port)
}
