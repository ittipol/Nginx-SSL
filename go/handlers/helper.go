package handlers

import (
	"go-nginx-ssl/errs"
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func getVarType(any interface{}) reflect.Type {
	return reflect.TypeOf(any)
}

func handleSuccess(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func handleSuccessWithMessage(c *fiber.Ctx, message string) error {
	c.Status(http.StatusOK)
	return c.JSON(fiber.Map{
		"message": message,
	})
}

func handleSuccessWithPayload(c *fiber.Ctx, payload interface{}) error {
	c.Status(http.StatusOK)
	return c.JSON(payload)
}

func handleError(c *fiber.Ctx, err error) error {
	if appError, errType := errs.ParseError(err); errType == errs.CustomerError {

		c.Status(appError.Code)
		return c.JSON(fiber.Map{
			"message": appError.Message,
		})

	}

	c.Status(http.StatusInternalServerError)
	return c.JSON(fiber.Map{
		"message": "unexpected error",
	})
}
