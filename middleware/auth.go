package middleware

import (
	"strings"
	"test_marketplace/models"
	"test_marketplace/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthRequired(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Status:  "error",
			Message: "Missing authorization header",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Status:  "error",
			Message: "Invalid token",
		})
	}
	c.Locals("user", claims)
	return c.Next()
}

func MerchantOnly(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)
	if claims.Role == "customer" {
		return c.Status(fiber.StatusForbidden).JSON(models.Response{
			Status:  "error",
			Message: "Access denied. Customer Only",
		})
	}
	return c.Next()
}

func CustomerOnly(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)
	if claims.Role != "customer" {
		return c.Status(fiber.StatusForbidden).JSON(models.Response{
			Status:  "error",
			Message: "Access denied. Customer only.",
		})
	}
	return c.Next()
}
