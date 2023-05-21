package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var allowedOrigins = []string{
	"http://abc:3000",
}

func CORS(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	origin, ok := headers["Origin"]

	if matched := contains(allowedOrigins, origin); ok && matched {
		fmt.Printf("Origin: %v \n\n", origin)
		c.Response().Header.Add("Access-Control-Allow-Origin", origin)
	}

	c.Response().Header.Add("Access-Control-Allow-Credentials", "false")
	c.Response().Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Response().Header.Add("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")

	return c.Next()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
