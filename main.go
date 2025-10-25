package main

import (
	"log"
	"test_marketplace/database"
	"test_marketplace/handlers"
	"test_marketplace/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()

	app := fiber.New()
	
	app.Use(logger.New())

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	app.Get("/products", handlers.GetProducts)

	api := app.Group("/api", middleware.AuthRequired)

	merchant := api.Group("/merchant", middleware.MerchantOnly)
	merchant.Post("/products", handlers.CreateProduct)
	merchant.Put("/products/:id", handlers.UpdateProduct)
	merchant.Delete("/products/:id", handlers.DeleteProduct)
	merchant.Get("/transactions", handlers.GetMerchantSales)
	merchant.Get("/transactions/report", handlers.GetSalesReport)
	merchant.Get("/transactions/customer/:customer_id", handlers.GetCustomerTransactions)
	merchant.Get("/transactions/top-products", handlers.GetTopProducts)
	merchant.Get("/transactions/:id", handlers.GetTransactionByID)

	customer := api.Group("/customer", middleware.CustomerOnly)
	customer.Post("/purchase", handlers.PurchaseProduct)
	customer.Get("/transactions", handlers.GetTransactionHistory)
	customer.Get("/transactions/:id", handlers.GetTransactionByID)

	api.Get("/transactions/:id", handlers.GetTransactionByID)

	log.Fatal(app.Listen(":3000"))
}
