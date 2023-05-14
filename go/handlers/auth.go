package handlers

import "github.com/gofiber/fiber/v2"

type authRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=1"`
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}
