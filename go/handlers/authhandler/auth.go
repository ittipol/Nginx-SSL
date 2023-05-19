package authhandler

import "github.com/gofiber/fiber/v2"

type loginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
	Verify(c *fiber.Ctx) error
}
