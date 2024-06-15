package server

import (
	"log"

	"github.com/Vajid-Hussain/machine-test/pkg/api/handler"
	"github.com/Vajid-Hussain/machine-test/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app *fiber.App
}

func InitServer(admin *handler.AdminHandler, user *handler.UserHandler, job *handler.JobHandler) *Server {
	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	routes.AdminRoutes(app.Group("/admin"), admin, job)
	routes.UserRoutes(app.Group("/"), user, job)

	return &Server{app: app}
}

func (s *Server) Start(port string) {
	err := s.app.Listen(port)
	log.Fatal("server starting error ", err)
}
