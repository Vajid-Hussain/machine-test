package handler

import (
	"net/http"

	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceUseCase "github.com/Vajid-Hussain/machine-test/pkg/usecase/interface"
	"github.com/Vajid-Hussain/machine-test/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	useCase interfaceUseCase.IUserUseCase
}

func NewUserHandler(usecase interfaceUseCase.IUserUseCase) *UserHandler {
	return &UserHandler{useCase: usecase}
}

func (u *UserHandler) UserSignup(ctx *fiber.Ctx) error {
	var userProfile requestmodels.UserSignup

	// Parse the incoming request body into variable.
	err := ctx.BodyParser(&userProfile)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(userProfile); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to sign up using the use case layer.
	res, err := u.useCase.UserSignUP(&userProfile)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(responsemodels.Responses(http.StatusCreated, "account created", res, nil))
}

func (u *UserHandler) UserLogin(ctx *fiber.Ctx) error {
	var req requestmodels.UserLogin

	// Parse the incoming request body into variable.
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	// Validate the parsed data.
	if validateErr := utils.Validator(req); len(validateErr) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, validateErr))
	}

	// Attempt to sign up using the use case layer.
	res, err := u.useCase.UserLogin(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "sucessfully Login", res, nil))
}
