package handlers

import (
	"log"
	"test_marketplace/database"
	"test_marketplace/models"
	"test_marketplace/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if input.Role != "merchant" && input.Role != "customer" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Role must be either 'merchant' or 'customer'",
		})
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	log.Printf("Registering user: %s, Role: %s", user.Email, user.Role)

	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Failed to create user: " + err.Error(),
		})
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to generate token",
		})
	}

	user.Password = ""

	return c.JSON(models.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data: models.LoginResponse{
			Token: token,
			User:  user,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	log.Printf("Login attempt for: %s", input.Email)

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		log.Printf("User not found: %s", input.Email)
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Status:  "error",
			Message: "Invalid credentials",
		})
	}

	log.Printf("User found: %s, Stored password: %s", user.Email, user.Password)

	if user.Password != input.Password {
		log.Printf("Password mismatch for user: %s", user.Email)
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Status:  "error",
			Message: "Invalid credentials",
		})
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to generate token",
		})
	}

	user.Password = ""

	log.Printf("Login successful for: %s", user.Email)

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Login successful",
		Data: models.LoginResponse{
			Token: token,
			User:  user,
		},
	})
}
