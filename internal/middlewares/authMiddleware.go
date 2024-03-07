package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	jwt_secret := os.Getenv("JWT_SECRET")
	cokkie := c.Cookies("token")
	if cokkie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized 1",
		})
	}

	// parse the token

	token, err := jwt.Parse(cokkie, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized 2",
		})
	}

	// check if token has expired

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized 3",
		})
	}
	/// check if the username in the token is the same as the requested username
	// Extract the username from the token
	username := claims["username"].(string)

	// Check if the requested username is provided in the query parameter
	requestedUsername := c.Query("username")

	// If not, check if it's provided in the JSON body
	if requestedUsername == "" {
		type RequestBody struct {
			Username string `json:"username"`
		}

		var requestBody RequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			return err
		}

		requestedUsername = requestBody.Username
	}

	if requestedUsername != username {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
