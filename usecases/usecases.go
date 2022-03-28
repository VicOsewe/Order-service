package usecases

import (
	"fmt"

	"github.com/VicOsewe/Order-service/application/interfaces"
	"github.com/VicOsewe/Order-service/domain"
	"github.com/VicOsewe/Order-service/repository"
)

type OrderService interface {
	CreateCustomer(customer *domain.Customer) (*domain.Customer, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	CreateOrder(order *domain.Order, orderProducts *[]domain.OrderProduct) (*string, error)
}

type Service struct {
	Repository repository.Repository
	SMS        interfaces.SMS
}

func NewOrderService(repo repository.Repository, sms interfaces.SMS) *Service {
	return &Service{
		Repository: repo,
		SMS:        sms,
	}
}

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
