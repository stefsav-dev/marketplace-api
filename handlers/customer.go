package handlers

import (
	"test_marketplace/database"
	"test_marketplace/models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	if err := database.DB.Preload("Merchant").Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to fetch products",
		})
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func PurchaseProduct(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)

	var req models.TransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	var product models.Product
	if err := database.DB.First(&product, req.ProductID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Status:  "error",
			Message: "Product not found",
		})
	}

	if product.Stock < req.Quantity {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Insufficient stock",
		})
	}

	totalPrice := product.Price * float64(req.Quantity)
	shippingCost := 10000.0
	discount := 0.0
	isFreeShipping := false

	if totalPrice > 15000 {
		isFreeShipping = true
		shippingCost = 0
	}

	if totalPrice > 50000 {
		discount = totalPrice * 0.1
	}

	finalPrice := totalPrice + shippingCost - discount

	transaction := models.Transaction{
		CustomerID:     claims.UserID,
		ProductID:      req.ProductID,
		Quantity:       req.Quantity,
		TotalPrice:     totalPrice,
		ShippingCost:   shippingCost,
		Discount:       discount,
		FinalPrice:     finalPrice,
		IsFreeShipping: isFreeShipping,
	}

	tx := database.DB.Begin()

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to create transaction",
		})
	}

	product.Stock -= req.Quantity
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to update product stock",
		})
	}

	tx.Commit()

	database.DB.Preload("Product").Preload("Product.Merchant").First(&transaction, transaction.ID)

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Purchase successful",
		Data:    transaction,
	})
}

func GetCustomerTransactions(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)

	var transactions []models.Transaction
	if err := database.DB.
		Preload("Product").
		Preload("Product.Merchant").
		Where("customer_id = ?", claims.UserID).
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
