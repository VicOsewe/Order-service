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

func TestOrderService_CreateProduct(t *testing.T) {

	type args struct {
		product *domain.Product
	}
	product := domain.Product{
		ID:        application.NewUUID(),
		Name:      gofakeit.CarModel(),
		UnitPrice: 300.0,
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Product
		wantErr bool
	}{
		{
			name: "happy case:",
			args: args{
				product: &product,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := postgres.NewOrderService()
			_, err := db.CreateProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestOrderService_CreateOrder(t *testing.T) {
	type args struct {
		order         *domain.Order
		orderProducts *[]domain.OrderProduct
	}

	db := postgres.NewOrderService()

	customer := domain.Customer{
		ID:          application.NewUUID(),
		PhoneNumber: gofakeit.PhoneFormatted(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		Password:    gofakeit.Password(true, true, true, true, true, 9),
	}
	cust, err := db.CreateCustomer(&customer)
	if err != nil {
		t.Fatalf("failed to create customer")
	}

	product := domain.Product{}
	prod, err := db.CreateProduct(&product)
	if err != nil {
		t.Fatalf("failed to create product")
	}
	order := domain.Order{
		TotalAmount: 3000,
		CustomerID:  cust.ID,
	}

	orderProduct := domain.OrderProduct{
		Product:         *prod,
		ProductQuantity: 1,
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Order
		wantErr bool
	}{
		{
			name: "happy case:",
			args: args{
				order: &order,
				orderProducts: &[]domain.OrderProduct{
					orderProduct,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := db.CreateOrder(tt.args.order, tt.args.orderProducts)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
