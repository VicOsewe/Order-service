package dao

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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
	Order       []Order `json:"-" gorm:"foreignKey:CustomerID"`
}

func (customer *Customer) BeforeCreate(tx *gorm.DB) error {
	// encrypt pin
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 8)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %v", err)
	}
	customer.Password = string(hashedPassword)
	return nil
}

//Order ...
type Order struct {
	ID           string         `json:"id"`
	TotalAmount  float64        `json:"total_amount"`
	CustomerID   string         `json:"customer_id"`
	OrderProduct []OrderProduct `json:"-"`
}

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unit_price"`
	Inventory int     `json:"inventory"`
}

type OrderProduct struct {
	ID              string  `json:"id"`
	OrderID         string  `json:"order_id"`
	ProductID       string  `json:"product_id"`
	ProductQuantity int     `json:"product_quantity"`
	Product         Product `json:"product" `
}

// OTP is used to persist and verify authorization codes
// (single use 'One Time PIN's)
type OTP struct {
	MSISDN            string    `json:"msisdn,omitempty" `
	Message           string    `json:"message,omitempty" `
	AuthorizationCode string    `json:"authorizationCode"`
	Timestamp         time.Time `json:"timestamp"`
	IsValid           bool      `json:"isValid"`
}
