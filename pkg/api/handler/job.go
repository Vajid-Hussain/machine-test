package handler

import (
	"errors"
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

// func (u I)


// ------------------------------------- Admin Job ----------------------------------------------


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
	if err != nil && !errors.Is(err, responsemodels.ErrNoActiveJob) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, err.Error(), res, nil))
}

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
