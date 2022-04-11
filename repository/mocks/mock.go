package mocks

import (
	"github.com/VicOsewe/Order-service/domain"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

// RepositoryMocks mocks the repository layer
type RepositoryMocks struct {
	MockCreateCustomer                   func(customer *domain.Customer) (*domain.Customer, error)
	MockCreateProduct                    func(product *domain.Product) (*domain.Product, error)
	MockCreateOrder                      func(order *domain.Order, orderProducts *[]domain.OrderProduct) (*domain.Order, error)
	MockGetCustomerByID                  func(customerID string) (*domain.Customer, error)
	MockGetProductByID                   func(productID string) (*domain.Product, error)
	MockGetCustomerByPhoneNumber         func(phoneNumber string) (*domain.Customer, error)
	MockGetProductByName                 func(name string) (*domain.Product, error)
	MockGetAllCustomerOrdersByCustomerID func(customerID string) (*[]domain.Order, error)
	MockGetAllProducts                   func() (*[]domain.Product, error)
	MockUpdateCustomer                   func(customer *domain.Customer) (*domain.Customer, error)
}

// NewRepositoryMocks inits a new instance of repository mocks with happy cases pre-defined
func NewRepositoryMocks() *RepositoryMocks {
	customerDetails := domain.Customer{
		ID:          uuid.New().String(),
		PhoneNumber: gofakeit.PhoneFormatted(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		Password:    gofakeit.Password(true, true, true, true, true, 9),
	}
	return &RepositoryMocks{
		MockCreateCustomer: func(customer *domain.Customer) (*domain.Customer, error) {
			return &customerDetails, nil
		},
		MockCreateProduct: func(product *domain.Product) (*domain.Product, error) {
			return nil, nil
		},
		MockCreateOrder: func(order *domain.Order, orderProducts *[]domain.OrderProduct) (*domain.Order, error) {
			return nil, nil
		},
		MockGetCustomerByID: func(customerID string) (*domain.Customer, error) {
			return nil, nil
		},
		MockGetProductByID: func(productID string) (*domain.Product, error) {
			return nil, nil
		},
		MockGetCustomerByPhoneNumber: func(phoneNumber string) (*domain.Customer, error) {
			return nil, nil
		},
		MockGetProductByName: func(name string) (*domain.Product, error) {
			return nil, nil
		},
		MockGetAllCustomerOrdersByCustomerID: func(customerID string) (*[]domain.Order, error) {
			return nil, nil
		},
		MockGetAllProducts: func() (*[]domain.Product, error) {
			return nil, nil
		},
		MockUpdateCustomer: func(customer *domain.Customer) (*domain.Customer, error) {
			return nil, nil
		},
	}
}

func (r *RepositoryMocks) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	return r.MockCreateCustomer(customer)
}
func (r *RepositoryMocks) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return r.MockCreateProduct(product)
}
func (r *RepositoryMocks) CreateOrder(order *domain.Order, orderProducts *[]domain.OrderProduct) (*domain.Order, error) {
	return r.MockCreateOrder(order, orderProducts)
}
func (r *RepositoryMocks) GetCustomerByID(customerID string) (*domain.Customer, error) {
	return r.MockGetCustomerByID(customerID)
}
func (r *RepositoryMocks) GetProductByID(productID string) (*domain.Product, error) {
	return r.MockGetProductByID(productID)
}
func (r *RepositoryMocks) GetCustomerByPhoneNumber(phoneNumber string) (*domain.Customer, error) {
	return r.MockGetCustomerByPhoneNumber(phoneNumber)
}
func (r *RepositoryMocks) GetProductByName(name string) (*domain.Product, error) {
	return r.MockGetProductByName(name)
}
func (r *RepositoryMocks) GetAllCustomerOrdersByCustomerID(customerID string) (*[]domain.Order, error) {
	return r.MockGetAllCustomerOrdersByCustomerID(customerID)
}
func (r *RepositoryMocks) GetAllProducts() (*[]domain.Product, error) {
	return r.MockGetAllProducts()
}
func (r *RepositoryMocks) UpdateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	return r.MockUpdateCustomer(customer)
}
