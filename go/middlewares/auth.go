package middlewares

import (
	"fmt"
	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	value, err := helpers.GetHeader(headers, "Authorization")

	if err != nil {
		return fiber.ErrUnauthorized
	}

	tokenString, err := helpers.GetBearerToken(value)

	if err != nil {
		return fiber.ErrUnauthorized
	}

	appJwtToken := appUtils.NewJwtUtil()
	token, err := appJwtToken.Validate(tokenString)

	if err != nil {
		return fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fiber.ErrUnauthorized
	}

	c.Set("Userid", fmt.Sprintf("%v", claims["id"]))
	return c.Next()
}
