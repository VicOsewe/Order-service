package domain

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;unique"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"user_name"`
	DateOfBirth string    `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	Password    string    `json:"password"`
	Order       []Order   `gorm:"foreignKey:CustomerID"`
}

type Order struct {
	ID           uuid.UUID      `json:"id"`
	TotalAmount  float64        `json:"total_amount"`
	CustomerID   uuid.UUID      `json:"customer_id"`
	OrderProduct []OrderProduct `json:"order_product"`
}

type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	UnitPrice float64   `json:"unit_price"`
}

type OrderProduct struct {
	ID              uuid.UUID `json:"id"`
	OrderID         uuid.UUID `json:"order_id"`
	ProductID       uuid.UUID
	ProductQuantity int     `json:"product_quantity"`
	Product         Product `json:"product" `
}
