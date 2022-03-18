package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

}

func SetUpRouter() {
	router := mux.NewRouter()

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
