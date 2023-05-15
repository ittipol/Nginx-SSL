package handlers

import (
	"go-nginx-ssl/apps"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/services"
	"strings"

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

	var req loginRequest

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

func (obj authHandler) Refresh(c *fiber.Ctx) error {
	// Get refresh token from Header
	// To prevent CSRF do not get refresh token from cookie
	headers := c.GetReqHeaders()

	if auth, ok := headers["Authorization"]; ok {

		// Bearer ALoCKD3FFOebeC2e3cYA4mLLb//2kCvkZziTLhVpI1TWnR4ZYkI3lak=
		parts := strings.Split(auth, " ")

		jwt := parts[1]

		_ = jwt

	}

	// return new access token and refresh token

	return c.JSON("ok")

}
