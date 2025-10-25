package handlers

import (
	"test_marketplace/database"
	"test_marketplace/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetTransactionHistory(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)

	var transactions []models.Transaction
	if err := database.DB.
		Preload("Product").
		Preload("Product.Merchant").
		Where("customer_id = ?", claims.UserID).
		Order("created_at DESC").
		Find(&transactions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to fetch transaction history",
		})
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Transaction history retrieved successfully",
		Data:    transactions,
	})
}

func GetMerchantSales(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)

	var sales []models.Transaction
	if err := database.DB.
		Preload("Customer").
		Preload("Product").
		Joins("JOIN products ON transactions.product_id = products.id").
		Where("products.merchant_id = ?", claims.UserID).
		Order("transactions.created_at DESC").
		Find(&sales).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to fetch sales data",
		})
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Sales data retrieved successfully",
		Data:    sales,
	})
}

func GetTransactionByID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)
	transactionID := c.Params("id")

	var transaction models.Transaction
	if err := database.DB.
		Preload("Product").
		Preload("Product.Merchant").
		Preload("Customer").
		First(&transaction, transactionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Status:  "error",
			Message: "Transaction not found",
		})
	}

	if claims.Role == "customer" && transaction.CustomerID != claims.UserID {
		return c.Status(fiber.StatusForbidden).JSON(models.Response{
			Status:  "error",
			Message: "Access denied to this transaction",
		})
	}

	if claims.Role == "merchant" {
		var product models.Product
		if err := database.DB.First(&product, transaction.ProductID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(models.Response{
				Status:  "error",
				Message: "Product not found",
			})
		}
		if product.MerchantID != claims.UserID {
			return c.Status(fiber.StatusForbidden).JSON(models.Response{
				Status:  "error",
				Message: "Access denied to this transaction",
			})
		}
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Transaction details retrieved successfully",
		Data:    transaction,
	})
}

func GetSalesReport(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)

	var request struct {
		StartDate string `query:"start_date"`
		EndDate   string `query:"end_date"`
	}

	if err := c.QueryParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: "Invalid query parameters",
		})
	}

	query := database.DB.
		Preload("Customer").
		Preload("Product").
		Joins("JOIN products ON transactions.product_id = products.id").
		Where("products.merchant_id = ?", claims.UserID)

	if request.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", request.StartDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Status:  "error",
				Message: "Invalid start date format. Use YYYY-MM-DD",
			})
		}
		query = query.Where("transactions.created_at >= ?", startDate)
	}

	if request.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", request.EndDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Status:  "error",
				Message: "Invalid end date format. Use YYYY-MM-DD",
			})
		}

		endDate = endDate.Add(24 * time.Hour)
		query = query.Where("transactions.created_at < ?", endDate)
	}

	var sales []models.Transaction
	if err := query.Order("transactions.created_at DESC").Find(&sales).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to generate sales report",
		})
	}

	var totalSales float64
	var totalTransactions int
	var totalDiscount float64
	var totalShippingCost float64

	for _, sale := range sales {
		totalSales += sale.FinalPrice
		totalTransactions++
		totalDiscount += sale.Discount
		totalShippingCost += sale.ShippingCost
	}

	report := struct {
		Sales   []models.Transaction   `json:"sales"`
		Summary map[string]interface{} `json:"summary"`
	}{
		Sales: sales,
		Summary: map[string]interface{}{
			"total_sales":         totalSales,
			"total_transactions":  totalTransactions,
			"total_discount":      totalDiscount,
			"total_shipping_cost": totalShippingCost,
			"average_order_value": totalSales / float64(totalTransactions),
		},
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Sales report generated successfully",
		Data:    report,
	})
}

func GetTopProducts(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JWTClaims)

	var topProducts []struct {
		ProductID    uint    `json:"product_id"`
		ProductName  string  `json:"product_name"`
		TotalSold    int     `json:"total_sold"`
		TotalRevenue float64 `json:"total_revenue"`
	}

	if err := database.DB.
		Table("transactions").
		Select(`
            products.id as product_id,
            products.name as product_name,
            SUM(transactions.quantity) as total_sold,
            SUM(transactions.final_price) as total_revenue
        `).
		Joins("JOIN products ON transactions.product_id = products.id").
		Where("products.merchant_id = ?", claims.UserID).
		Group("products.id, products.name").
		Order("total_sold DESC").
		Limit(10).
		Find(&topProducts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: "Failed to fetch top products",
		})
	}

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Top products retrieved successfully",
		Data:    topProducts,
	})
}
