package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Role      string    `gorm:"size:20;not null" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	MerchantID  uint      `gorm:"not null" json:"merchant_id"`
	Merchant    User      `gorm:"foreignKey:MerchantID" json:"merchant"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int       `gorm:"not null" json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	CustomerID     uint      `gorm:"not null" json:"customer_id"`
	Customer       User      `gorm:"foreignKey:CustomerID" json:"customer"`
	ProductID      uint      `gorm:"not null" json:"product_id"`
	Product        Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity       int       `gorm:"not null" json:"quantity"`
	TotalPrice     float64   `gorm:"type:decimal(10,2);not null" json:"total_price"`
	ShippingCost   float64   `gorm:"type:decimal(10,2);not null" json:"shipping_cost"`
	Discount       float64   `gorm:"type:decimal(10,2);not null" json:"discount"`
	FinalPrice     float64   `gorm:"type:decimal(10,2);not null" json:"final_price"`
	IsFreeShipping bool      `gorm:"not null" json:"is_free_shipping"`
	CreatedAt      time.Time `json:"created_at"`
}
type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
