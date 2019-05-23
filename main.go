package main

import (
	"./handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//http request handler developed using mux package
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/posts", handler.GetPosts).Methods("GET")
	subRouter.HandleFunc("/post/create", handler.CreatePost).Methods("POST")
	subRouter.HandleFunc("/post/{id}", handler.GetPost).Methods("GET")
	subRouter.HandleFunc("/companies", handler.GetCompanies).Methods("GET")
	subRouter.HandleFunc("/users", handler.GetUsers).Methods("GET")
	subRouter.HandleFunc("/user/create", handler.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":3001", router))
}

//program starts from here
func main() {
	handleRequests()
}
