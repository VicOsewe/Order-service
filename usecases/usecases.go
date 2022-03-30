package usecases

import (
	"fmt"

	"github.com/VicOsewe/Order-service/application/interfaces"
	"github.com/VicOsewe/Order-service/domain"
	"github.com/VicOsewe/Order-service/repository"
)

//OrderService ...
type OrderService interface {
	CreateCustomer(customer *domain.Customer) (*domain.Customer, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	CreateOrder(order *domain.Order, orderProducts *[]domain.OrderProduct) (*string, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*domain.Customer, error)
	GetProductByName(name string) (*domain.Product, error)
	GetAllCustomerOrdersByPhoneNumber(phoneNumber string) (*[]domain.Order, error)
	GetAllProducts() (*[]domain.Product, error)
}

//Service ...
type Service struct {
	Repository repository.Repository
	SMS        interfaces.SMS
}

//NewOrderService ...
func NewOrderService(repo repository.Repository, sms interfaces.SMS) *Service {
	return &Service{
		Repository: repo,
		SMS:        sms,
	}
}

//CreateCustomer creates a customer record in the database
func (s *Service) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {

	if customer == nil {
		return nil, fmt.Errorf("customer fields required")
	}

	cust, err := s.Repository.CreateCustomer(customer)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer :%v", err)
	}

	return cust, nil
}

//CreateProduct creates a product record in the database
func (s *Service) CreateProduct(product *domain.Product) (*domain.Product, error) {

	if product == nil {
		return nil, fmt.Errorf("product fields required")
	}

	prod, err := s.Repository.CreateProduct(product)
	if err != nil {
		return nil, fmt.Errorf("failed to create product :%v", err)
	}

	return prod, nil
}

//CreateOrder creates order record in the database
func (s *Service) CreateOrder(order *domain.Order, orderProducts *[]domain.OrderProduct) (*string, error) {
	if order.CustomerID == "" {
		return nil, fmt.Errorf("ensure customer_id is provided")
	}
	//ensure that the customer exists before create an order for them
	customer, err := s.Repository.GetCustomerByID(order.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve customer records with err of %v", err)
	}
	if customer == nil {
		return nil, fmt.Errorf("failed to find a customer with id %v", order.CustomerID)
	}

	for _, orderProduct := range *orderProducts {
		if orderProduct.ProductID == "" {
			return nil, fmt.Errorf("ensure that a product(s) is provided")
		}
		if orderProduct.ProductQuantity == 0 {
			return nil, fmt.Errorf("ensure that product quantity is more than one")
		}
		// ensure that the product exists in the database
		product, err := s.Repository.GetProductByID(orderProduct.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to retreive product with err of %v", err)
		}
		if product == nil {
			return nil, fmt.Errorf("failed to find product with id %v", product.ID)
		}
	}

	ord, err := s.Repository.CreateOrder(order, orderProducts)
	if err != nil {
		return nil, fmt.Errorf("failed to create an order :%v", err)
	}
	message := fmt.Sprintf("Dear customer you order has been created with an order id of %v, our team will contact you shortly to finalize it", ord.ID)
	err = s.SMS.SendSMS(message, customer.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to send sms to customer with an error of %v", err)
	}

	return &ord.ID, nil
}

//GetCustomerByPhoneNumber fetches cutomer records using a customers phone number
func (s *Service) GetCustomerByPhoneNumber(phoneNumber string) (*domain.Customer, error) {
	customer, err := s.Repository.GetCustomerByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer record with phone_number %v and err %v", phoneNumber, err)
	}
	return customer, nil

}

//GetProductByName fetches product record using the product name
func (s *Service) GetProductByName(name string) (*domain.Product, error) {
	product, err := s.Repository.GetProductByName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product with name:  %v and err: %v ", name, err)
	}
	return product, nil
}

//GetAllCustomerOrdersByPhoneNumber fetches all orders for a given customer
func (s *Service) GetAllCustomerOrdersByPhoneNumber(phoneNumber string) (*[]domain.Order, error) {
	customer, err := s.Repository.GetCustomerByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer record with phone_number %v and err %v", phoneNumber, err)
	}
	if customer == nil {
		return nil, fmt.Errorf("failed to get customer")
	}

	order, err := s.Repository.GetAllCustomerOrdersByCustomerID(customer.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve orders of the cutomer with id: %v and err of %v", customer.ID, err)
	}
	return order, nil
}

//GetAllProducts gets all the products in the database
func (s *Service) GetAllProducts() (*[]domain.Product, error) {
	products, err := s.Repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to get products with an error of %v", err)
	}
	if products == nil {
		return nil, fmt.Errorf("failed to find any products")
	}
	return products, nil
}
