package handlers

import (
	"go-nginx-ssl/apps"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/services"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService services.AuthService
	validate    apps.Validator
}

func NewAuthHandler(authService services.AuthService, validate apps.Validator) AuthHandler {
	return &authHandler{authService, validate}
}

func (obj authHandler) Login(c *fiber.Ctx) error {

	var req authRequest

	if err := c.BodyParser(&req); err != nil {
		logs.Error(err)
		return handleError(c, err)
	}

	if err := obj.validate.ValidatePayload(req); err != nil {
		logs.Error(err)
		return handleError(c, err)
	}

	res, err := obj.authService.Login(req.Email, req.Password)
	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}

	return handleSuccessWithPayload(c, res)
}
