package dao

import "gorm.io/gorm"

//Customer ...
type Customer struct {
	ID          string  `json:"id" gorm:"primaryKey;unique"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Username    string  `json:"user_name"`
	DateOfBirth string  `json:"date_of_birth"`
	Gender      string  `json:"gender"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Address     string  `json:"address"`
	Password    string  `json:"password"`
	Order       []Order `gorm:"foreignKey:CustomerID"`
}

func (customer *Customer) BeforeCreate(tx *gorm.DB) error {
	// encrypt pin
	return nil
}

//Order ...
type Order struct {
	ID           string         `json:"id"`
	TotalAmount  float64        `json:"total_amount"`
	CustomerID   string         `json:"customer_id"`
	OrderProduct []OrderProduct `json:"order_product"`
}

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unit_price"`
	Inventory int
}

type OrderProduct struct {
	ID              string  `json:"id"`
	OrderID         string  `json:"order_id"`
	ProductID       string  `json:"product_id"`
	ProductQuantity int     `json:"product_quantity"`
	Product         Product `json:"product" `
}
