package repository

import "github.com/VicOsewe/Order-service/domain/dao"

type Repository interface {
	CreateCustomer(customer *dao.Customer) (*dao.Customer, error)
	CreateProduct(product *dao.Product) (*dao.Product, error)
	CreateOrder(order *dao.Order, orderProducts *[]dao.OrderProduct) (*dao.Order, error)
	GetCustomerByID(customerID string) (*dao.Customer, error)
	GetProductByID(productID string) (*dao.Product, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*dao.Customer, error)
	GetProductByName(name string) (*dao.Product, error)
	GetAllCustomerOrdersByCustomerID(customerID string) (*[]dao.Order, error)
	GetAllProducts() (*[]dao.Product, error)
	UpdateCustomer(customer *dao.Customer) (*dao.Customer, error)
	SaveOTP(otp dao.OTP) error
}
