package postgres_test

import (
	"testing"

	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/VicOsewe/Order-service/infrastucture/databases/postgres"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

func TestOrderService_CreateCustomer(t *testing.T) {

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
		order         *dao.Order
		orderProducts *[]dao.OrderProduct
	}

	db := postgres.NewOrderService()

	customer := dao.Customer{
		ID:          uuid.New().String(),
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

	product := dao.Product{}
	prod, err := db.CreateProduct(&product)
	if err != nil {
		t.Fatalf("failed to create product")
	}
	order := dao.Order{
		TotalAmount: 3000,
		CustomerID:  cust.ID,
	}

	orderProduct := dao.OrderProduct{
		Product:         *prod,
		ProductQuantity: 1,
	}

	tests := []struct {
		name    string
		args    args
		want    *dao.Order
		wantErr bool
	}{
		{
			name: "happy case:",
			args: args{
				order: &order,
				orderProducts: &[]dao.OrderProduct{
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

func TestOrderService_UpdateCustomer(t *testing.T) {

	type args struct {
		customer *dao.Customer
	}
	db := postgres.NewOrderService()

	customer := dao.Customer{
		ID:          uuid.New().String(),
		PhoneNumber: gofakeit.PhoneFormatted(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		Password:    gofakeit.Password(true, true, true, true, true, 9),
	}
	customerDetails, err := db.CreateCustomer(&customer)
	if err != nil {
		t.Fatalf("failed to create customer")
	}

	tests := []struct {
		name    string
		args    args
		want    *dao.Customer
		wantErr bool
	}{
		{
			name: "happy case:",
			args: args{
				customer: &dao.Customer{
					PhoneNumber: customerDetails.PhoneNumber,
					FirstName:   "new name",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.UpdateCustomer(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.UpdateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.PhoneNumber != customer.PhoneNumber {
				t.Errorf("OrderService.UpdateCustomer() expected phone number to be: %v but got %v", customer.PhoneNumber, got.PhoneNumber)
				return
			}
		})
	}
}
