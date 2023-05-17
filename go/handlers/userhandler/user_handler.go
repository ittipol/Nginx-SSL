package userhandler

import (
	"go-nginx-ssl/handlers"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/services/usersrv"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService usersrv.UserService
}

func NewUserHandler(userService usersrv.UserService) UserHandler {
	return &userHandler{userService}
}

func (obj userHandler) Register(c *fiber.Ctx) error {

	var req registerRequest

	if err := c.BodyParser(&req); err != nil {
		logs.Error(err)
		return handlers.HandleError(c, err)
	}

	err := obj.userService.Register(req.Email, req.Password, req.Name)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	return handlers.HandleSuccess(c)
}
