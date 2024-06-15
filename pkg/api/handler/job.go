package handler

import (
	"net/http"

	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceUseCase "github.com/Vajid-Hussain/machine-test/pkg/usecase/interface"
	"github.com/Vajid-Hussain/machine-test/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type JobHandler struct {
	useCase interfaceUseCase.IJobUseCase
}

func NewJobHandler(useCase interfaceUseCase.IJobUseCase) *JobHandler {
	return &JobHandler{useCase: useCase}
}

// @Summary Upload and decode resume
// @Description Upload a resume file and decode its content
// @Tags user job
// @Accept  multipart/form-data
// @Produce  json
// @Param resume formData file true "Resume file"
// @Success 200 {object} responsemodels.Response "Resume successfully added"
// @Failure 400 {object} responsemodels.Response "Invalid input"
// @Security authorization
// @Router /resume [post]
func (u *JobHandler) DecodeResume(ctx *fiber.Ctx) error {
	var (
		req requestmodels.Resume
		err error
	)

	req.UserID = ctx.Locals("userID").(string)
	req.Resume, err = ctx.FormFile("resume")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	res, err := u.useCase.DecodeResume(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "resume succesfully added", res, nil))
}

// @Summary Apply for a job
// @Description Submit a job application
// @Tags apply job user
// @Accept  json
// @Produce  json
// @Param jobApplication body requestmodels.JobApplication true "Job Application"
// @Success 200 {object} responsemodels.Response "Successfully applied for the job"
// @Failure 400 {object} responsemodels.Response "Invalid input"
// @Security authorization
// @Router /applied [post]
func (u *JobHandler) ApplyJob(ctx *fiber.Ctx) error {
	var req requestmodels.JobApplication
	req.UserID = ctx.Locals("userID").(string)

	// Parse the incoming request body into variable.
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(req); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to apply jobs using the use case layer.
	res, err := u.useCase.ApplyJob(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "job added", res, nil))
}

// @Summary Get applied jobs
// @Description Retrieve a list of applied jobs with pagination and optional search
// @Tags apply job user
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(5)
// @Param search query string false "Search keyword"
// @Success 200 {object} responsemodels.Response "Successfully retrieved applied jobs"
// @Failure 400 {object} responsemodels.Response "Invalid input"
// @Security authorization
// @Router /applied [get]
func (u *JobHandler) GetAppliedJob(ctx *fiber.Ctx) error {
	var (
		req        requestmodels.GetAppliedJob
		pagination requestmodels.Pagination
	)
	req.UserID = ctx.Locals("userID").(string)

	// Parse the incoming request body into variable.
	err := ctx.QueryParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}
	err = ctx.QueryParser(&pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(req); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to get applied jobs using the use case layer.
	res, err := u.useCase.GetAppliedJob(&req, &pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "", res, nil))
}

// ------------------------------------- Admin Job ----------------------------------------------

// @Summary Create a job
// @Description Create a new job posting
// @Tags Job Management Admin
// @Accept  json
// @Produce  json
// @Param job body requestmodels.CreateJob true "Job Create"
// @Success 201 {object} responsemodels.Response "Successfully created job"
// @Failure 400 {object} responsemodels.Response "Invalid input"
// @Security authorization
// @Router /admin/job [post]
func (u *JobHandler) CreateJob(ctx *fiber.Ctx) error {
	var req requestmodels.CreateJob
	req.PostedBy = ctx.Locals("adminID").(string)

	// Parse the incoming request body into variable.
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(req); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to Create jobs using the use case layer.
	res, err := u.useCase.CreateJob(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "job added", res, nil))
}

// @Summary Get jobs
// @Description Retrieve a list of jobs with pagination and optional search
// @Tags user job
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(5)
// @Param search query string false "Search keyword"
// @Success 200 {object} responsemodels.Response "Successfully retrieved jobs"
// @Failure 400 {object} responsemodels.Response "Invalid input"
// @Security authorization
// @Router /job [get]
func (u *JobHandler) GetJob(ctx *fiber.Ctx) error {
	var (
		req        requestmodels.JobSearch
		pagination requestmodels.Pagination
	)

	// Parse the incoming request body into variable.
	err := ctx.QueryParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}
	err = ctx.QueryParser(&pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(req); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to get jobs using the use case layer.
	res, err := u.useCase.GetJob(&req, &pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "", res, nil))
}

// @Summary Get jobs
// @Description Retrieve a list of jobs with pagination and optional search
// @Tags Job Management Admin
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(5)
// @Param search query string false "Search keyword"
// @Success 200 {object} responsemodels.Response "Successfully retrieved jobs"
// @Failure 400 {object} responsemodels.Response "Invalid input"
// @Security authorization
// @Router /admin/job [get]
func (u *JobHandler) GetJobAdmin(ctx *fiber.Ctx) error {
	var (
		req        requestmodels.JobSearch
		pagination requestmodels.Pagination
	)

	// Parse the incoming request body into variable.
	err := ctx.QueryParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}
	err = ctx.QueryParser(&pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(req); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to get jobs using the use case layer.
	res, err := u.useCase.GetJob(&req, &pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "", res, nil))
}

// @Summary Delete a job
// @Description Delete an existing job posting
// @Tags Job Management Admin
// @Accept  json
// @Produce  json
// @Param job body requestmodels.DeleteJob true "Job Delete"
// @Success 200 {object} responsemodels.Response "Successfully deleted job"
// @Failure 400 {object} responsemodels.Response "Invalid input"
// @Security authorization
// @Router /admin/job [delete]
func (u *JobHandler) DeleteJob(ctx *fiber.Ctx) error {
	var req requestmodels.DeleteJob

	// Parse the incoming request body into variable.
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(req); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to delete jobs using the use case layer.
	err = u.useCase.DeleteJob(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "job deleted succsfully", "", nil))
}

// @Summary Get complete details of job
// @Description Retrieve details of a specific job with pagination
// @Tags Job Management Admin
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(5)
// @Param jobid query int true "Job ID"
// @Success 200 {object} responsemodels.Response "Successfully retrieved job details"
// @Failure 400 {object} responsemodels.Response "Invalid input"
// @Security authorization
// @Router /admin/job/details [get]
func (u *JobHandler) GetJobDetails(ctx *fiber.Ctx) error {
	var (
		req        requestmodels.JobID
		pagination requestmodels.Pagination
	)

	// Parse the incoming request body into variable.
	err := ctx.QueryParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}
	err = ctx.QueryParser(&pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(req); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to get applied jobs using the use case layer.
	res, err := u.useCase.GetJobDetails(&req, &pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "", res, nil))
}
