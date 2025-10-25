package handlers

import (
	"test_marketplace/database"
	"test_marketplace/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)

	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	product.MerchantID = claims.UserID

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Failed to create product : " + err.Error(),
		})
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Product created successfully",
		Data:    product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)
	productID := c.Params("id")

	var product models.Product
	if err := database.DB.Where("id = ? AND merchant_id = ?", productID, claims.UserID).First(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Status:  "error",
			Message: "Product not found",
		})
	}

	var updateData models.Product
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if updateData.Name != "" {
		product.Name = updateData.Name
	}
	if updateData.Description != "" {
		product.Description = updateData.Description
	}
	if updateData.Price > 0 {
		product.Price = updateData.Price
	}
	if updateData.Stock >= 0 {
		product.Stock = updateData.Stock
	}

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Failed to update product : " + err.Error(),
		})
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Product updated successfully",
		Data:    product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)
	productID := c.Params("id")

	var product models.Product
	if err := database.DB.Where("id = ? AND merchant_id = ?", productID, claims.UserID).First(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Status:  "error",
			Message: "Product not found",
		})
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to delete product : " + err.Error(),
		})
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Product deleted successfully",
	})
}

func GetMerchantTransactions(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)

	var transactions []models.Transaction
	if err := database.DB.
		Preload("Customer").
		Preload("Product").
		Joins("JOIN products ON transactions.product_id = products.id").
		Where("products.merchant_id = ?", claims.UserID).
		Find(&transactions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to fetch transactions",
		})
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Transactions retrieved successfully",
		Data:    transactions,
	})
}
