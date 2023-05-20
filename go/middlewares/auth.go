package middlewares

import (
	"fmt"
	"go-nginx-ssl/appUtils"
	"go-nginx-ssl/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func AuthorizeJWT(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()
	fmt.Printf("Headers: %v\n\n", headers)
	value, err := helpers.GetHeader(headers, "Authorization")

	if err != nil {
		return fiber.ErrUnauthorized
	}

	tokenString, err := helpers.GetBearerToken(value)

	if err != nil {
		return fiber.ErrUnauthorized
	}

	appJwt := appUtils.NewJwtUtil(
		[]byte(viper.GetString("app.jwt_access_token_secret")),
		[]byte(viper.GetString("app.jwt_refresh_token_secret")),
	)
	token, err := appJwt.Validate(tokenString, appUtils.AccessTokenSecretKey)

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
