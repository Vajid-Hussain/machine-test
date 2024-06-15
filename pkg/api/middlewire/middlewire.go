package middlewire

import (
	"net/http"
	"strings"

	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	"github.com/Vajid-Hussain/machine-test/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

var token config.JWTConfig

func NewMiddewire(tokenSecret config.JWTConfig) {
	token = tokenSecret
}

func VerifyAdminToken(ctx *fiber.Ctx) error {

	authHeader := ctx.Get("authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(responsemodels.Responses(http.StatusUnauthorized, "Unauthorized: Missing or invalid authorization token", nil, nil))
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	adminID, err := utils.VerifyAcessToken(accessToken, token.SecretKeyAdmin)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(responsemodels.Responses(http.StatusUnauthorized, "", nil, err.Error()))
	}

	ctx.Locals("adminID", adminID)

	return ctx.Next()
}

func VerifyUserToken(ctx *fiber.Ctx) error {

	authHeader := ctx.Get("authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(responsemodels.Responses(http.StatusUnauthorized, "Unauthorized: Missing or invalid authorization token", nil, nil))
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	userID, err := utils.VerifyAcessToken(accessToken, token.SecretKeyUser)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(responsemodels.Responses(http.StatusUnauthorized, "", nil, err.Error()))
	}

	// set user id in context
	ctx.Locals("userID", userID)

	return ctx.Next()
}
