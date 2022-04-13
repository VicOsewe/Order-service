package usecases

import (
	"fmt"
	"testing"

	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/VicOsewe/Order-service/interfaces/repository/mocks"
	smsMocks "github.com/VicOsewe/Order-service/interfaces/services/mocks"
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
		customer *dao.Customer
	}

	customer := dao.Customer{
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
		want    *dao.Customer
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
				repo.MockGetCustomerByPhoneNumber = func(phoneNumber string) (*dao.Customer, error) {
					return nil, fmt.Errorf("failed to get customer details")
				}
			}
			if tt.name == "sad case:create customer details failed" {
				repo.MockCreateCustomer = func(customer *dao.Customer) (*dao.Customer, error) {
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
				return
			}
		})
	}
}

func TestService_CreateProduct(t *testing.T) {

	type args struct {
		product *dao.Product
	}
	product := dao.Product{
		ID:        uuid.New().String(),
		Name:      gofakeit.CarModel(),
		UnitPrice: 300.0,
	}

	tests := []struct {
		name    string
		args    args
		want    *dao.Product
		wantErr bool
	}{
		{
			name: "happy case:product record created",
			args: args{
				product: &product,
			},
			wantErr: false,
		},
		{
			name: "sad case:get product failure",
			args: args{
				product: &product,
			},
			wantErr: true,
		},
		{
			name: "sad case:create product in database failure",
			args: args{
				product: &product,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepositoryMocks()
			sms := smsMocks.NewSMSMocks()
			s := NewOrderService(repo, sms)
			if tt.name == "sad case:get product failure" {
				repo.MockGetProductByName = func(name string) (*dao.Product, error) {
					return nil, fmt.Errorf("failed to get product")
				}
			}
			if tt.name == "sad case:create product in database failure" {
				repo.MockGetProductByName = func(name string) (*dao.Product, error) {
					product := dao.Product{}
					return &product, nil
				}
				repo.MockCreateProduct = func(product *dao.Product) (*dao.Product, error) {
					return nil, fmt.Errorf("failed to create product")
				}
			}
			got, err := s.CreateProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.wantErr == false {
				t.Errorf("Service.CreateProduct() failed to create customer")
				return
			}
		})
	}
}
