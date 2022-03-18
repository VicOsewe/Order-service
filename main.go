package main

import (
	"log"
	"net/http"
	"os"

	"github.com/VicOsewe/Order-service/infrastucture/databases/postgres"
	"github.com/gorilla/mux"
)

func main() {
	SetUpRouter()

}

func SetUpRouter() {
	router := mux.NewRouter()

	_, err := postgres.InitializeDatabase()
	if err != nil {
		log.Printf("failed to connect to database :%v", err)
		os.Exit(1)
	}

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
