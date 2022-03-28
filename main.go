package main

import (
	"net/http"

	"github.com/VicOsewe/Order-service/infrastucture/databases/postgres"
	"github.com/VicOsewe/Order-service/presentation/rest"
	"github.com/VicOsewe/Order-service/usecases"

	"github.com/gorilla/mux"
)

func main() {
	SetUpRouter()

}

func SetUpRouter() {
	router := mux.NewRouter()
	rep := postgres.NewOrderService()
	usecases := usecases.NewOrderService(rep)
	handler := rest.NewHandler(usecases)

	router.Path("/customers").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.CreateCustomer)
	router.Path("/products").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.CreateProduct)
	router.Path("/orders").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.CreateOrder)
	router.Path("/products").Methods(http.MethodOptions, http.MethodGet).HandlerFunc(handler.GetAllProducts)
	router.Path("/customers").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.GetCustomerByPhoneNumber)
	router.Path("/product").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.GetProductByName)
	router.Path("/orders").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.GetAllCustomerOrdersByPhoneNumber)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
