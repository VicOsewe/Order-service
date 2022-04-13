package usecases

import (
	"fmt"

	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/VicOsewe/Order-service/interfaces/repository"
	interfaces "github.com/VicOsewe/Order-service/interfaces/services"
)

type OrderService interface {
	CreateCustomer(customer *dao.Customer) (*dao.Customer, error)
	CreateProduct(product *dao.Product) (*dao.Product, error)
	CreateOrder(order *dao.Order, orderProducts *[]dao.OrderProduct) (*string, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*dao.Customer, error)
	GetProductByName(name string) (*dao.Product, error)
	GetAllCustomerOrdersByPhoneNumber(phoneNumber string) (*[]dao.Order, error)
	GetAllProducts() (*[]dao.Product, error)
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

//CreateCustomer checks to see if a customer records exists and if not it creates a new record
// of the customer details
func (s *Service) CreateCustomer(customer *dao.Customer) (*dao.Customer, error) {
	if customer == nil {
		return nil, fmt.Errorf("customer fields required")
	}
	customerDetails, err := s.Repository.GetCustomerByPhoneNumber(customer.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if customerDetails.ID == "" {
		cust, err := s.Repository.CreateCustomer(customer)
		if err != nil {
			return nil, fmt.Errorf("failed to create customer :%v", err)
		}
		return cust, nil
	}

	return customerDetails, nil
}

//CreateProduct checks to see if a product record exists and if not it creates a new record
func (s *Service) CreateProduct(product *dao.Product) (*dao.Product, error) {
	if product == nil {
		return nil, fmt.Errorf("product fields required")
	}

	productDetails, err := s.Repository.GetProductByName(product.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create products")
	}

	if productDetails.ID == "" {
		prod, err := s.Repository.CreateProduct(product)
		if err != nil {
			return nil, fmt.Errorf("failed to create product :%v", err)
		}
		return prod, nil
	}

	return productDetails, nil
}

func (s *Service) CreateOrder(order *dao.Order, orderProducts *[]dao.OrderProduct) (*string, error) {
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

func (s *Service) GetCustomerByPhoneNumber(phoneNumber string) (*dao.Customer, error) {
	customer, err := s.Repository.GetCustomerByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer record with phone_number %v and err %v", phoneNumber, err)
	}
	return customer, nil

}

func (s *Service) GetProductByName(name string) (*dao.Product, error) {
	product, err := s.Repository.GetProductByName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product with name:  %v and err: %v ", name, err)
	}
	return product, nil
}

func (s *Service) GetAllCustomerOrdersByPhoneNumber(phoneNumber string) (*[]dao.Order, error) {
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

func (s *Service) GetAllProducts() (*[]dao.Product, error) {
	products, err := s.Repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to get products with an error of %v", err)
	}
	if products == nil {
		return nil, fmt.Errorf("failed to find any products")
	}
	return products, nil
}
