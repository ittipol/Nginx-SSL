package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {

	app := fiber.New()

	group := app.Group("", logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		// time.Sleep(time.Second * 1)
		return c.JSON("OK")
	})

	group.Post("/user/:id", func(c *fiber.Ctx) error {

		payload := struct {
			ID int `json:"id"`
		}{}

		id := utils.CopyString(c.Params("id"))

		fmt.Printf("Param ID: %v\n", id)

		if err := c.BodyParser(&payload); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status": fiber.ErrBadRequest,
			})
		}

		fmt.Printf("Request Body: %v \n", payload)

		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"payload": payload,
		})
	})

	log.Fatal(app.Listen(":5000"))

}
