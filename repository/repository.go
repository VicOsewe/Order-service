package repository

import "github.com/VicOsewe/Order-service/domain"

type Repository interface {
	CreateCustomer(customer *domain.Customer) (*domain.Customer, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	CreateOrder(order *domain.Order, orderProducts *[]domain.OrderProduct) (*domain.Order, error)
	GetCustomerByID(customerID string) (*domain.Customer, error)
	GetProductByID(productID string) (*domain.Product, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*domain.Customer, error)
	GetProductByName(name string) (*domain.Product, error)
	GetAllCustomerOrdersByCustomerID(customerID string) (*[]domain.Order, error)
	GetAllProducts() (*[]domain.Product, error)
}
