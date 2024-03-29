package usecases

import (
	"fmt"
	"time"

	"github.com/VicOsewe/Order-service/application"
	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/VicOsewe/Order-service/domain/dto"

	"github.com/VicOsewe/Order-service/interfaces/repository"
	interfaces "github.com/VicOsewe/Order-service/interfaces/services"
)

type OrderService interface {
	VerifyPhoneNumber(phoneNumber string) (*dto.OtpResponse, error)
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

// VerifyPhoneNumber checks validity of a phone number by sending an OTP to it
func (s *Service) VerifyPhoneNumber(phoneNumber string) (*dto.OtpResponse, error) {
	if phoneNumber == "" {
		return nil, fmt.Errorf("phone number field is required")
	}

	customerDetails, err := s.Repository.GetCustomerByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}

	if customerDetails.PhoneNumber != "" {
		return nil, fmt.Errorf("the phone number already exists")
	}

	code, err := s.GenerateAndSendOTP(phoneNumber)
	if err != nil {
		return nil, err
	}
	otpResponse := &dto.OtpResponse{
		OTP: code,
	}

	return otpResponse, nil
}

//GenerateAndSendOTP generates an OTP, persists it in the database and sends it to the
//supplied phone number as a text message
func (s *Service) GenerateAndSendOTP(phoneNumber string) (string, error) {
	code, err := application.GenerateOTP()
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf(" Your One Time Password is %s", code)
	otp := dao.OTP{
		MSISDN:            phoneNumber,
		Message:           message,
		AuthorizationCode: code,
		Timestamp:         time.Now(),
		IsValid:           true,
	}

	err = s.Repository.SaveOTP(otp)
	if err != nil {
		return "", err
	}
	err = s.SMS.SendSMS(message, phoneNumber)
	if err != nil {
		return "", err
	}

	return code, nil
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
		if orderProduct.ProductQuantity == 0 {
			return nil, fmt.Errorf("ensure that product quantity is more than one")
		}
		//fetch the product form the database
		product, err := s.Repository.GetProductByName(orderProduct.Product.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to retreive product with err of %v", err)
		}
		if product == nil {
			return nil, fmt.Errorf("product with name %v does not exist", product.Name)
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
