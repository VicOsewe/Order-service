package mocks

import (
	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

// RepositoryMocks mocks the repository layer
type RepositoryMocks struct {
	MockCreateCustomer                   func(customer *dao.Customer) (*dao.Customer, error)
	MockCreateProduct                    func(product *dao.Product) (*dao.Product, error)
	MockCreateOrder                      func(order *dao.Order, orderProducts *[]dao.OrderProduct) (*dao.Order, error)
	MockGetCustomerByID                  func(customerID string) (*dao.Customer, error)
	MockGetProductByID                   func(productID string) (*dao.Product, error)
	MockGetCustomerByPhoneNumber         func(phoneNumber string) (*dao.Customer, error)
	MockGetProductByName                 func(name string) (*dao.Product, error)
	MockGetAllCustomerOrdersByCustomerID func(customerID string) (*[]dao.Order, error)
	MockGetAllProducts                   func() (*[]dao.Product, error)
	MockUpdateCustomer                   func(customer *dao.Customer) (*dao.Customer, error)
}

// NewRepositoryMocks inits a new instance of repository mocks with happy cases pre-defined
func NewRepositoryMocks() *RepositoryMocks {
	customerDetails := dao.Customer{
		ID:          uuid.New().String(),
		PhoneNumber: gofakeit.PhoneFormatted(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		Password:    gofakeit.Password(true, true, true, true, true, 9),
	}
	product := dao.Product{
		ID:        uuid.New().String(),
		Name:      gofakeit.BeerName(),
		UnitPrice: gofakeit.Price(300.0, 400.0),
		Inventory: 100,
	}
	return &RepositoryMocks{
		MockCreateCustomer: func(customer *dao.Customer) (*dao.Customer, error) {
			return &customerDetails, nil
		},
		MockCreateProduct: func(product *dao.Product) (*dao.Product, error) {
			return nil, nil
		},
		MockCreateOrder: func(order *dao.Order, orderProducts *[]dao.OrderProduct) (*dao.Order, error) {
			return nil, nil
		},
		MockGetCustomerByID: func(customerID string) (*dao.Customer, error) {
			return &customerDetails, nil
		},
		MockGetProductByID: func(productID string) (*dao.Product, error) {
			return &product, nil
		},
		MockGetCustomerByPhoneNumber: func(phoneNumber string) (*dao.Customer, error) {
			return &customerDetails, nil
		},
		MockGetProductByName: func(name string) (*dao.Product, error) {
			return &product, nil
		},
		MockGetAllCustomerOrdersByCustomerID: func(customerID string) (*[]dao.Order, error) {
			return nil, nil
		},
		MockGetAllProducts: func() (*[]dao.Product, error) {
			return nil, nil
		},
		MockUpdateCustomer: func(customer *dao.Customer) (*dao.Customer, error) {
			return nil, nil
		},
	}
}

func (r *RepositoryMocks) CreateCustomer(customer *dao.Customer) (*dao.Customer, error) {
	return r.MockCreateCustomer(customer)
}
func (r *RepositoryMocks) CreateProduct(product *dao.Product) (*dao.Product, error) {
	return r.MockCreateProduct(product)
}
func (r *RepositoryMocks) CreateOrder(order *dao.Order, orderProducts *[]dao.OrderProduct) (*dao.Order, error) {
	return r.MockCreateOrder(order, orderProducts)
}
func (r *RepositoryMocks) GetCustomerByID(customerID string) (*dao.Customer, error) {
	return r.MockGetCustomerByID(customerID)
}
func (r *RepositoryMocks) GetProductByID(productID string) (*dao.Product, error) {
	return r.MockGetProductByID(productID)
}
func (r *RepositoryMocks) GetCustomerByPhoneNumber(phoneNumber string) (*dao.Customer, error) {
	return r.MockGetCustomerByPhoneNumber(phoneNumber)
}
func (r *RepositoryMocks) GetProductByName(name string) (*dao.Product, error) {
	return r.MockGetProductByName(name)
}
func (r *RepositoryMocks) GetAllCustomerOrdersByCustomerID(customerID string) (*[]dao.Order, error) {
	return r.MockGetAllCustomerOrdersByCustomerID(customerID)
}
func (r *RepositoryMocks) GetAllProducts() (*[]dao.Product, error) {
	return r.MockGetAllProducts()
}
func (r *RepositoryMocks) UpdateCustomer(customer *dao.Customer) (*dao.Customer, error) {
	return r.MockUpdateCustomer(customer)
}
