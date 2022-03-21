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

	router.Path("/customer").Methods(http.MethodOptions, http.MethodPost).HandlerFunc(handler.CreateCustomer)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
