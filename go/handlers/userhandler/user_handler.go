package userhandler

import (
	"fmt"
	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/handlers"
	"go-nginx-ssl/logs"
	"go-nginx-ssl/services/usersrv"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService usersrv.UserService
	validate    appUtils.ValidatorUtil
}

func NewUserHandler(userService usersrv.UserService, validate appUtils.ValidatorUtil) UserHandler {
	return &userHandler{userService, validate}
}

func (obj userHandler) Register(c *fiber.Ctx) error {

	var req registerRequest

	if err := c.BodyParser(&req); err != nil {
		logs.Error(err)
		return handlers.HandleError(c, err)
	}

	logs.Info(fmt.Sprintf("%v \n", req))

	if err := obj.validate.ValidatePayload(req); err != nil {
		logs.Error(err)
		return handlers.HandleError(c, err)
	}

	err := obj.userService.Register(req.Email, req.Password, req.Name)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	return handlers.HandleSuccess(c)
}
