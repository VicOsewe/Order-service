package domain

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Password    string
}
