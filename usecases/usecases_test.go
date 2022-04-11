package usecases

import (
	"fmt"
	"testing"

	smsMocks "github.com/VicOsewe/Order-service/application/interfaces/mocks"
	"github.com/VicOsewe/Order-service/domain"
	"github.com/VicOsewe/Order-service/repository/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

func MockNewOrderService() *Service {
	repo := mocks.NewRepositoryMocks()
	sms := smsMocks.NewSMSMocks()
	return &Service{
		Repository: repo,
		SMS:        sms,
	}
}

func TestService_CreateCustomer(t *testing.T) {
	type args struct {
		customer *domain.Customer
	}

	customer := domain.Customer{
		ID:          uuid.New().String(),
		PhoneNumber: gofakeit.PhoneFormatted(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		Password:    gofakeit.Password(true, true, true, true, true, 9),
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Customer
		wantErr bool
	}{
		{
			name: "happy case:customer record created",
			args: args{
				customer: &customer,
			},
			wantErr: false,
		},
		{
			name:    "sad case:customer input is nil",
			args:    args{},
			wantErr: true,
		},
		{
			name: "sad case:get customer details failed",
			args: args{
				customer: &customer,
			},
			wantErr: true,
		},
		{
			name: "sad case:create customer details failed",
			args: args{
				customer: &customer,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepositoryMocks()
			sms := smsMocks.NewSMSMocks()
			s := NewOrderService(repo, sms)
			if tt.name == "sad case:get customer details failed" {
				repo.MockGetCustomerByPhoneNumber = func(phoneNumber string) (*domain.Customer, error) {
					return nil, fmt.Errorf("failed to get customer details")
				}
			}
			if tt.name == "sad case:create customer details failed" {
				repo.MockCreateCustomer = func(customer *domain.Customer) (*domain.Customer, error) {
					return nil, fmt.Errorf("failed to create customer")
				}
			}

			got, err := s.CreateCustomer(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.wantErr == false {
				t.Errorf("Service.CreateCustomer() failed to create customer")
			}
		})
	}
}
