package authhandler

import (
	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/handlers"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/services/authsrv"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService authsrv.AuthService
	validate    appUtils.ValidatorUtil
}

func NewAuthHandler(authService authsrv.AuthService, validate appUtils.ValidatorUtil) AuthHandler {
	return &authHandler{authService, validate}
}

func (obj authHandler) Login(c *fiber.Ctx) error {

	var req loginRequest

	if err := c.BodyParser(&req); err != nil {
		logs.Error(err)
		return handlers.HandleError(c, err)
	}

	if err := obj.validate.ValidatePayload(req); err != nil {
		logs.Error(err)
		return handlers.HandleError(c, err)
	}

	res, err := obj.authService.Login(req.Email, req.Password)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	c.Cookie(&fiber.Cookie{
		Name:        "refresh-token",
		Value:       res.RefreshToken,
		SessionOnly: true,
		HTTPOnly:    true,
		Secure:      true,
		SameSite:    "lax", // default is lax mode
	})

	return handlers.HandleSuccessWithPayload(c, res)
}

func (obj authHandler) Refresh(c *fiber.Ctx) error {
	// Get refresh token from Header
	// To prevent CSRF do not get refresh token from cookie
	headers := c.GetReqHeaders()

	res, err := obj.authService.Refresh(headers)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	// c.Cookie(&fiber.Cookie{
	// 	Name:        "refresh-token",
	// 	Value:       res.RefreshToken,
	// 	SessionOnly: true,
	// 	HTTPOnly:    true,
	// 	Secure:      true,
	// 	SameSite:    "strict", // default is lax mode
	// })

	return handlers.HandleSuccessWithPayload(c, res)
}

func (obj authHandler) Verify(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	err := obj.authService.Verify(headers)

	if err != nil {
		return handlers.HandleError(c, err)
	}

	return handlers.HandleSuccessWithMessage(c, "Valid Token")
}
