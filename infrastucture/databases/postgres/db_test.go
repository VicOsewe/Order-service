package postgres_test

import (
	"testing"

	"github.com/VicOsewe/Order-service/application"
	"github.com/VicOsewe/Order-service/domain"
	"github.com/VicOsewe/Order-service/infrastucture/databases/postgres"
	"github.com/brianvoe/gofakeit"
)

func TestOrderService_CreateCustomer(t *testing.T) {

	type args struct {
		customer *domain.Customer
	}
	customer := domain.Customer{
		ID:          application.NewUUID(),
		PhoneNumber: gofakeit.PhoneFormatted(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		Password:    gofakeit.Password(true, true, true, true, true, 9),
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy case:",
			args: args{
				customer: &customer,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := postgres.NewOrderService()
			got, err := db.CreateCustomer(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.FirstName != tt.args.customer.FirstName {
				t.Errorf("OrderService.CreateCustomer() expected firstname to be = %v, but got %v", tt.args.customer.FirstName, got.FirstName)
			}
		})
	}
}
