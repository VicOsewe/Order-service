package usecases_test

import (
	"testing"

	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/VicOsewe/Order-service/infrastucture/databases/postgres"
	ait "github.com/VicOsewe/Order-service/infrastucture/services/AIT"
	"github.com/VicOsewe/Order-service/usecases"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

func TestService_CreateCustomer(t *testing.T) {
	type args struct {
		customer *dao.Customer
	}

	newCustomerRecord := dao.Customer{
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
				customer: &newCustomerRecord,
			},
			wantErr: false,
		},
		{
			name:    "sad case:customer input is nil",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := postgres.NewOrderService()
			smsService := ait.NewAITService()

			s := usecases.NewOrderService(db, smsService)
			got, err := s.CreateCustomer(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "happy case:customer record created" {
				//ensure that the customer record exists in the db
				cust, err := s.Repository.GetCustomerByID(got.ID)
				if err != nil {
					t.Fatalf("failed to get customer")
				}
				if cust == nil {
					t.Errorf("Service.CreateCustomer() expected customer record to exist")
					return
				}
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
			name:    "sad case:product is nil",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := postgres.NewOrderService()
			smsService := ait.NewAITService()

			s := usecases.NewOrderService(db, smsService)

			got, err := s.CreateProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "happy case:product record created" {
				//ensure that the product exists in the db
				prod, err := s.Repository.GetProductByID(got.ID)
				if err != nil {
					t.Fatalf("failed to get customer")
				}
				if prod == nil {
					t.Errorf("Service.CreateProduct()) expected product record to exist")
					return
				}
			}
			if got == nil && tt.wantErr == false {
				t.Errorf("Service.CreateProduct() failed to create customer")
				return
			}
		})
	}
}
