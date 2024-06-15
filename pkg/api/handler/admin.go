package handler

import (
	"net/http"

	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceUseCase "github.com/Vajid-Hussain/machine-test/pkg/usecase/interface"
	"github.com/Vajid-Hussain/machine-test/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	useCase interfaceUseCase.IAdminUseCase
}

func NewAdminHandler(usecase interfaceUseCase.IAdminUseCase) *AdminHandler {
	return &AdminHandler{useCase: usecase}
}

func (u *AdminHandler) AdminLogin(ctx *fiber.Ctx) error {
	var req requestmodels.AdminLogin

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
	res, err := u.useCase.AdminLogin(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(responsemodels.Responses(http.StatusOK, "sucessfully Login", res, nil))
}
