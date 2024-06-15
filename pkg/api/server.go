package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func InitServer() *Server {
	app := fiber.New()

	return &Server{app: app}
}

func (s *Server) Start(port string) {
	err := s.app.Listen(port)
	log.Fatal("server starting error ", err)
}
