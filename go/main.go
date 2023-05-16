package main

import (
	"fmt"

	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/handlers"
	"go-nginx-ssl/middlewares"
	"go-nginx-ssl/repositories"
	"go-nginx-ssl/services"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"gorm.io/gorm"
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

	var db *gorm.DB

	appValidator := appUtils.NewValidatorUtil()
	appJwt := appUtils.NewJwtUtil()

	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository, appJwt)
	authHandler := handlers.NewAuthHandler(authService, appValidator)

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		// Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	// app.Post("/user/:id", func(c *fiber.Ctx) error {

	// 	payload := struct {
	// 		ID int `json:"id"`
	// 	}{}

	// 	id := utils.CopyString(c.Params("id"))

	// 	fmt.Printf("Param ID: %v\n", id)

	// 	if err := c.BodyParser(&payload); err != nil {
	// 		c.Status(fiber.StatusBadRequest)
	// 		return c.JSON(fiber.Map{
	// 			"status": fiber.ErrBadRequest,
	// 		})
	// 	}

	// 	fmt.Printf("Request Body: %v \n", payload)

	// 	c.Status(fiber.StatusOK)
	// 	return c.JSON(fiber.Map{
	// 		"status":  fiber.StatusOK,
	// 		"payload": payload,
	// 	})
	// })

	app.Post("/auth", authHandler.Login)
	app.Post("/refresh", authHandler.Refresh)
	// app.Post("/verify", authHandler.Verify)

	app.Use("/health", middlewares.AuthorizeJWT)
	app.Get("/health", func(c *fiber.Ctx) error {

		headers := c.GetRespHeaders()

		fmt.Printf("%v\n", headers["Userid"])

		return c.JSON("OK")
	})

	api := app.Group("/api", middlewares.AuthorizeJWT)
	api.Get("/test", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusOK, "OK...")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port"))))

}
