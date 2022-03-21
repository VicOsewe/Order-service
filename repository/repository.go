package repository

import "github.com/VicOsewe/Order-service/domain"

type Repository interface {
	CreateCustomer(customer *domain.Customer) (*domain.Customer, error)
}
