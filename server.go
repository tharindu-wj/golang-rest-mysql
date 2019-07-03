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
	subRouter.HandleFunc("/post/delete/{id}", handler.RemovePost).Methods("DELETE")

	subRouter.HandleFunc("/companies", handler.GetCompanies).Methods("GET")

	subRouter.HandleFunc("/categories", handler.GetCategories).Methods("GET")
	subRouter.HandleFunc("/category/{id}", handler.GetCategory).Methods("GET")

	subRouter.HandleFunc("/users", handler.GetUsers).Methods("GET")
	subRouter.HandleFunc("/user/create", handler.CreateUser).Methods("POST")
	subRouter.HandleFunc("/user/{id}", handler.GetUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}

//program starts from here
func main() {
	handleRequests()
}
