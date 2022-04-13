package rest

import (
	"log"
	"net/http"

	"github.com/VicOsewe/Order-service/application"
	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/VicOsewe/Order-service/domain/dto"
	"github.com/VicOsewe/Order-service/usecases"
)

type HandlersInterfaces interface {
	CreateCustomer() http.HandlerFunc
	CreateProduct() http.HandlerFunc
	CreateOrder(w http.ResponseWriter, r *http.Request)
	GetCustomerByPhoneNumber(phoneNumber string) (*dao.Customer, error)
	GetProductByName(name string) (*dao.Product, error)
	GetAllCustomerOrdersByPhoneNumber(phoneNumber string) (*[]dao.Order, error)
	GetAllProducts() (*[]dao.Product, error)
}

type HandlersImplementation struct {
	auth struct {
		username string
		password string
	}
	Usecases usecases.OrderService
}

func NewHandler(usecases usecases.OrderService) HandlersImplementation {

	app := HandlersImplementation{
		Usecases: usecases,
	}
	app.auth.password = application.GetEnv("AUTH_PASSWORD")
	app.auth.username = application.GetEnv("AUTH_USERNAME")
	if app.auth.username == "" {
		log.Fatal("basic auth username must be provided")
	}

	if app.auth.password == "" {
		log.Fatal("basic auth password must be provided")
	}
	return app
}

//CreateCustomer creates a record of customer details
func (h *HandlersImplementation) CreateCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customer := dao.Customer{}
		err := UnmarshalJSONToStruct(w, r, &customer)
		if err != nil {
			response := dto.APIFailureResponse{
				Error: err.Error(),
				APIResponse: dto.APIResponse{
					Message: "failed to unmarshal to struct",
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
					Message: "failed to validate customer info",
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
					Message: "failed to normalize phone number",
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
					Message: "failed to create customer record",
				},
			}
			HandlerResponse(w, http.StatusBadRequest, response)
			return

		}
		response := dto.APIResponse{
			Message: "customer created successfully",
			Body:    cust,
		}

		HandlerResponse(w, http.StatusCreated, response)
	}

}

func (h *HandlersImplementation) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		product := dao.Product{}
		err := UnmarshalJSONToStruct(w, r, &product)
		if err != nil {
			response := dto.APIFailureResponse{
				Error: err.Error(),
				APIResponse: dto.APIResponse{
					Message: "failed to unmarshal to struct",
				},
			}
			HandlerResponse(w, http.StatusInternalServerError, response)
			return

		}
		if product.Name == "" || product.UnitPrice == 0 {
			response := dto.APIFailureResponse{
				Error: "failed to validate product details",
				APIResponse: dto.APIResponse{

					Message: "invalid request data, ensure name, inventory and unit_price is provided",
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
					Message: "failed to create product record",
				},
			}
			HandlerResponse(w, http.StatusBadRequest, response)
			return

		}
		response := dto.APIResponse{
			Message: "product created successfully",
			Body:    prod,
		}

		HandlerResponse(w, http.StatusAccepted, response)
	}

}

func (h *HandlersImplementation) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := dto.OrderInput{}
	err := UnmarshalJSONToStruct(w, r, &order)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "failed to unmarshal to struct",
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
				Message: "failed to create order record",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	response := dto.APIResponse{
		Message: "order created successfully",
		Body:    prod,
	}

	HandlerResponse(w, http.StatusAccepted, response)

}

func (h *HandlersImplementation) GetCustomerByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	customer := dao.Customer{}
	err := UnmarshalJSONToStruct(w, r, &customer)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "failed to unmarshal to struct",
			},
		}
		HandlerResponse(w, http.StatusInternalServerError, response)
		return

	}
	if customer.PhoneNumber == "" {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "invalid request data, phone_number is provided",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return
	}

	prod, err := h.Usecases.GetCustomerByPhoneNumber(customer.PhoneNumber)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "failed to fetch customer record",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	response := dto.APIResponse{
		Message: "customer retrieved successfully",
		Body:    prod,
	}

	HandlerResponse(w, http.StatusAccepted, response)

}

func (h *HandlersImplementation) GetProductByName(w http.ResponseWriter, r *http.Request) {
	product := dao.Product{}
	err := UnmarshalJSONToStruct(w, r, &product)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "failed to unmarshal to struct",
			},
		}
		HandlerResponse(w, http.StatusInternalServerError, response)
		return

	}
	if product.Name == "" {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "invalid request data, name is provided",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return
	}

	prod, err := h.Usecases.GetProductByName(product.Name)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "failed to fetch product record",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	response := dto.APIResponse{
		Message: "product retrieved successfully",
		Body:    prod,
	}

	HandlerResponse(w, http.StatusAccepted, response)

}

func (h *HandlersImplementation) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	prod, err := h.Usecases.GetAllProducts()
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "failed to fetch products record",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	response := dto.APIResponse{
		Message: "products retrieved successfully",
		Body:    prod,
	}

	HandlerResponse(w, http.StatusAccepted, response)

}

func (h *HandlersImplementation) GetAllCustomerOrdersByPhoneNumber(w http.ResponseWriter, r *http.Request) {

	customer := dao.Customer{}
	err := UnmarshalJSONToStruct(w, r, &customer)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "failed to unmarshal to struct",
			},
		}
		HandlerResponse(w, http.StatusInternalServerError, response)
		return

	}
	if customer.PhoneNumber == "" {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "invalid request data, phone_number is provided",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return
	}

	prod, err := h.Usecases.GetAllCustomerOrdersByPhoneNumber(customer.PhoneNumber)
	if err != nil {
		response := dto.APIFailureResponse{
			Error: err.Error(),
			APIResponse: dto.APIResponse{
				Message: "failed to fetch order record",
			},
		}
		HandlerResponse(w, http.StatusBadRequest, response)
		return

	}
	response := dto.APIResponse{
		Message: "orders retrieved successfully",
		Body:    prod,
	}

	HandlerResponse(w, http.StatusAccepted, response)

}
