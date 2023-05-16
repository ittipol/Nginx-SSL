package handlers

import (
	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/services"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService services.AuthService
	validate    appUtils.ValidatorUtil
}

func NewAuthHandler(authService services.AuthService, validate appUtils.ValidatorUtil) AuthHandler {
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
	// headers := c.GetReqHeaders()

	// return new access token and refresh token
	return c.JSON("ok")

}

func (obj authHandler) Verify(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	err := obj.authService.Verify(headers)

	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}

	return handleSuccessWithMessage(c, "Valid Token")
}
