package usecases

import (
	"fmt"

	"github.com/VicOsewe/Order-service/domain"
	"github.com/VicOsewe/Order-service/repository"
)

type OrderService interface {
	CreateCustomer(customer *domain.Customer) (*domain.Customer, error)
}

type Service struct {
	Repository repository.Repository
}

func NewOrderService(repo repository.Repository) *Service {
	return &Service{
		Repository: repo,
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
