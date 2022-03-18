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

	//connect to db
	err := postgres.SetUpDB()
	if err != nil {
		log.Printf("failed to connect to database :%v", err)
		os.Exit(1)
	}

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
