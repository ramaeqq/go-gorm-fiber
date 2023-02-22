package middleware

import (
	"ramaeqq/go-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(m *fiber.Ctx) error {

	token := m.Get("x-token")

	if token == "" {
		return m.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "unauthenticated",
		})
	}

	// _, err := utils.VerifyToken(token)
	claims, err := utils.DecodeToken(token)
	if err != nil {
		return m.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return m.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"Message": "forbidden access",
		})
	}

	m.Locals("userInfo", claims)
	// m.Locals("role", claims["role"])

	return m.Next()

}

func PermisionAdd(m *fiber.Ctx) error {
	return m.Next()
}
