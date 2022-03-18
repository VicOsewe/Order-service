package main

import (
	"net/http"

	"github.com/VicOsewe/Order-service/infrastucture/databases/postgres"
	"github.com/gorilla/mux"
)

func main() {
	SetUpRouter()

}

func SetUpRouter() {
	router := mux.NewRouter()

	_ = postgres.InitializeDatabase()

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
