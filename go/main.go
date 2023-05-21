package main

import (
	"fmt"
	"log"
	"time"

	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/database"
	"go-nginx-ssl/handlers/authhandler"
	"go-nginx-ssl/handlers/userhandler"
	"go-nginx-ssl/middlewares"
	"go-nginx-ssl/repositories"
	"go-nginx-ssl/services/authsrv"
	"go-nginx-ssl/services/usersrv"

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

	initTimeZone()
	db := initDbConnection()

	appValidator := appUtils.NewValidatorUtil()
	appJwt := appUtils.NewJwtUtil(
		[]byte(viper.GetString("app.jwt_access_token_secret")),
		[]byte(viper.GetString("app.jwt_refresh_token_secret")),
	)

	userRepository := repositories.NewUserRepository(db)
	authService := authsrv.NewAuthService(userRepository, appJwt)
	userService := usersrv.NewUserService(userRepository, appJwt)

	authHandler := authhandler.NewAuthHandler(authService, appValidator)
	userHandler := userhandler.NewUserHandler(userService, appValidator)

	// ========================================================================

	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://localhost",
	// 	AllowHeaders:     "Origin, Content-Type, Authorization",
	// 	AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	// 	AllowCredentials: false,
	// }))

	app.Use(middlewares.CORS, logger.New(logger.Config{
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

	app.Post("/login", authHandler.Login)

	app.Use("/token/refresh", middlewares.RefreshTokenAuthorizeJWT)
	app.Post("/token/refresh", authHandler.Refresh)

	app.Post("/verify", authHandler.Verify)

	app.Post("/register", userHandler.Register)

	app.Use("/user/profile", middlewares.AuthorizeJWT)
	app.Get("/user/profile", userHandler.Profile)

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

func initTimeZone() {
	// LoadLocation looks for the IANA Time Zone database
	// List of tz database time zones
	// https: //en.wikipedia.org/wiki/List_of_tz_database_time_zones
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	// init system time zone
	time.Local = location

	// timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	// fmt.Println(timeInUTC.In(location))
}

func initDbConnection() *gorm.DB {
	return database.GetDbConnection(
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.db"),
		false,
	)
}
