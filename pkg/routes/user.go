package routes

import (
	"github.com/Vajid-Hussain/machine-test/pkg/api/handler"
	"github.com/Vajid-Hussain/machine-test/pkg/api/middlewire"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(engin fiber.Router, user *handler.UserHandler, job *handler.JobHandler) {
	engin.Post("/signup", user.UserSignup)
	engin.Post("/login", user.UserLogin)

	engin.Use(middlewire.VerifyUserToken)

	engin.Post("/resume", job.DecodeResume)

	jobManagement := engin.Group("/job")
	{
		jobManagement.Get("/", job.GetJob)
	}
}
