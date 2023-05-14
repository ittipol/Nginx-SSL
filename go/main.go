package main

import (
	"fmt"
	"go-nginx-ssl/apps"
	"go-nginx-ssl/handlers"
	"go-nginx-ssl/services"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {

	appValidator := apps.Newvalidator(validator.New())
	authService := services.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService, appValidator)

	app := fiber.New()

	group := app.Group("", logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
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

	app.Post("/auth", authHandler.Login)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port"))))

}
