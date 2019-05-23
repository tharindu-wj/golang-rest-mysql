package main

import (
	"./handler"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//home page route
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

//http request handler developed using mux package
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("/", homePage)
	subRouter.HandleFunc("/posts", handler.GetPosts).Methods("GET")
	subRouter.HandleFunc("/companies", handler.GetCompanies).Methods("GET")
	subRouter.HandleFunc("/users", handler.GetUsers).Methods("GET")
	subRouter.HandleFunc("/user/{id}", handler.GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

//program starts from here
func main() {
	handleRequests()
}
