package userhandler

import "github.com/gofiber/fiber/v2"

type registerRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Name     string `validate:"required"`
}

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
}
