package rest

import (
	"net/http"

	"github.com/VicOsewe/Order-service/application/dto"
	"github.com/VicOsewe/Order-service/domain"
	"github.com/VicOsewe/Order-service/usecases"
)

type HandlersInterfaces interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type HandlersImplementation struct {
	Usecases usecases.OrderService
}

func NewHandler(usecases usecases.OrderService) HandlersImplementation {
	return HandlersImplementation{
		Usecases: usecases,
	}
}

func (h *HandlersImplementation) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := domain.Customer{}
	err := UnmarshalJSONToStruct(w, r, &customer)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to unmarshal to struct",
			},
		}
		HandlerResponse(w, http.StatusInternalServerError, response)
		return

	}
	err = ValidateCustomerInfo(customer)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "failed to validate customer info",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	normalizedPhoneNumber, err := NormalizePhoneNumber(customer.PhoneNumber)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "failed to normalize phone number",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	customer.PhoneNumber = normalizedPhoneNumber
	cust, err := h.Usecases.CreateCustomer(&customer)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "failed to create customer record",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	response := dto.APIResponse{
		Message:    "customer created successfully",
		Body:       cust,
		StatusCode: http.StatusCreated,
	}

	HandlerResponse(w, http.StatusAccepted, response)

}

func (h *HandlersImplementation) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := domain.Product{}
	err := UnmarshalJSONToStruct(w, r, &product)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to unmarshal to struct",
			},
		}
		HandlerResponse(w, http.StatusInternalServerError, response)
		return

	}
	if product.Name == "" || product.UnitPrice == 0 {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "invalid request data, ensure name and unit_price is provided",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return
	}

	prod, err := h.Usecases.CreateProduct(&product)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "failed to create product record",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	response := dto.APIResponse{
		Message:    "product created successfully",
		Body:       prod,
		StatusCode: http.StatusCreated,
	}

	HandlerResponse(w, http.StatusAccepted, response)

}

func (h *HandlersImplementation) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := dto.OrderInput{}
	err := UnmarshalJSONToStruct(w, r, &order)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to unmarshal to struct",
			},
		}
		HandlerResponse(w, http.StatusInternalServerError, response)
		return

	}

	prod, err := h.Usecases.CreateOrder(&order.Order, &order.OrderProduct)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "failed to create order record",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	response := dto.APIResponse{
		Message:    "order created successfully",
		Body:       prod,
		StatusCode: http.StatusCreated,
	}

	HandlerResponse(w, http.StatusAccepted, response)

}
