package main

import (
	"net/http"

	"github.com/VicOsewe/Order-service/infrastucture/databases/postgres"
	ait "github.com/VicOsewe/Order-service/infrastucture/services/AIT"

	"github.com/VicOsewe/Order-service/presentation/rest"
	"github.com/VicOsewe/Order-service/usecases"

	"github.com/gorilla/mux"
)

func main() {
	SetUpRouter()

}

func SetUpRouter() {
	router := mux.NewRouter()
	db := postgres.NewOrderService()
	smsService := ait.NewAITService()

	usecases := usecases.NewOrderService(db, smsService)
	handler := rest.NewHandler(usecases)
	router.Use(handler.BasicAuth())

	router.Path("/verify").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.VerifyCustomerPhoneNumber())
	router.Path("/customers").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.CreateCustomer())
	router.Path("/products").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.CreateProduct())
	router.Path("/orders").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.CreateOrder())
	router.Path("/products").Methods(http.MethodOptions, http.MethodGet).HandlerFunc(handler.GetAllProducts())
	router.Path("/customers/{phone_number}").Methods(http.MethodOptions, http.MethodGet).HandlerFunc(handler.GetCustomerByPhoneNumber())
	router.Path("/products/{name}").Methods(http.MethodOptions, http.MethodGet).HandlerFunc(handler.GetProductByName())
	router.Path("/orders").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.GetAllCustomerOrdersByPhoneNumber)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
