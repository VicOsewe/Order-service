package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/VicOsewe/Order-service/interfaces/repository/mocks"
	smsMocks "github.com/VicOsewe/Order-service/interfaces/services/mocks"
	"github.com/VicOsewe/Order-service/usecases"
	"github.com/brianvoe/gofakeit"
)

func TestHandlersImplementation_CreateCustomer(t *testing.T) {
	type args struct {
		url        string
		httpMethod string
		body       io.Reader
	}

	fakeRepo := mocks.NewRepositoryMocks()
	fakeSms := smsMocks.NewSMSMocks()
	usecases := usecases.NewOrderService(fakeRepo, fakeSms)
	h := NewHandler(usecases)

	customer := dao.Customer{
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		PhoneNumber: gofakeit.Phone(),
		Password:    gofakeit.Password(true, true, true, true, true, 9),
		Email:       gofakeit.Email(),
	}
	marshalledCustomer, err := json.Marshal(customer)
	if err != nil {
		t.Errorf("unable to marshal payload: %s", err)
		return
	}

	invalidCustomer := dao.Customer{
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		PhoneNumber: gofakeit.Phone(),
		Password:    gofakeit.Password(true, true, true, true, true, 9),
		Email:       gofakeit.Email(),
	}
	marshalledInvalidCustomer, err := json.Marshal(invalidCustomer)
	if err != nil {
		t.Errorf("unable to marshal payload: %s", err)
		return
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "happy case: create customer success",
			args: args{
				url:        fmt.Sprintf("%s/customers", "http://localhost:5000"),
				httpMethod: http.MethodPost,
				body:       bytes.NewBuffer(marshalledCustomer),
			},
			wantStatus: http.StatusCreated,
			wantErr:    false,
		},
		{
			name: "sad case: failed to create customer",
			args: args{
				url:        fmt.Sprintf("%s/customers", "http://localhost:5000"),
				httpMethod: http.MethodPost,
				body:       bytes.NewBuffer(marshalledInvalidCustomer),
			},
			wantStatus: http.StatusCreated,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "sad case:create customer details failed" {
				fakeRepo.MockCreateCustomer = func(customer *dao.Customer) (*dao.Customer, error) {
					return nil, fmt.Errorf("failed to create customer")
				}
			}
			req, err := http.NewRequest(tt.args.httpMethod, tt.args.url, tt.args.body)
			if err != nil {
				t.Errorf("can't create new request: %v", err)
				return
			}
			response := httptest.NewRecorder()

			svr := h.CreateCustomer()
			svr.ServeHTTP(response, req)

			if tt.wantStatus != response.Code {
				t.Errorf("expected status %d, got %d", tt.wantStatus, response.Code)
				return
			}
		})
	}
}
