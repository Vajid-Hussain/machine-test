package routes

import (
	"github.com/Vajid-Hussain/machine-test/pkg/api/handler"
	"github.com/Vajid-Hussain/machine-test/pkg/api/middlewire"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(engin fiber.Router, admin *handler.AdminHandler, job *handler.JobHandler) {
	engin.Post("/login", admin.AdminLogin)

	engin.Use(middlewire.VerifyAdminToken)

	jobManagement := engin.Group("/job")
	{
		jobManagement.Post("/", job.CreateJob)
		jobManagement.Get("/", job.GetJobAdmin)
		jobManagement.Delete("/", job.DeleteJob)
		jobManagement.Get("/details", job.GetJobDetails)
	}
}
