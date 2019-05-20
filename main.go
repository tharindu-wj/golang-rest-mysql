package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//single post object
type Post struct {
	Id      int
	Title   string
	Body    string
	Created string
	User    int
}

//single user object
type User struct {
	Id      int
	Name    string
	Email   string
	Created string
}

//single company object
type Company struct {
	Id       int
	Name     string
	Location string
}

//database connection
func dbConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "asdf1234"
	dbName := "go_leafycode"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}

//get all posts
func allPosts(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM posts ORDER BY created_at DESC")
	if err != nil {
		panic(err.Error())
	}
	post := Post{}
	posts := []Post{}
	for selDB.Next() {
		var id, user int
		var title, body, created_at string
		err = selDB.Scan(&id, &title, &body, &user, &created_at)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Title = title
		post.Body = body
		post.User = user
		post.Created = created_at
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)

	defer db.Close()
}

//get user data
func getUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM users WHERE id=?", userID)
	if err != nil {
		panic(err.Error())
	}

	user := User{}

	for selDB.Next() {
		var id int
		var name, email, created_at string
		err = selDB.Scan(&id, &name, &email, &created_at)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Name = name
		user.Email = email
		user.Created = created_at
	}
	defer db.Close()
	json.NewEncoder(w).Encode(user)
}

//get all companies
func allCompanies(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM companies")
	if err != nil {
		panic(err.Error())
	}
	company := Company{}
	companies := []Company{}
	for selDB.Next() {
		var id int
		var name, location string
		err = selDB.Scan(&id, &name, &location)
		if err != nil {
			panic(err.Error())
		}
		company.Id = id
		company.Name = name
		company.Location = location
		companies = append(companies, company)
	}
	defer db.Close()
	json.NewEncoder(w).Encode(companies)
}

//home page route
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

//http request handler developed using mux package
func handleRequests() {
	testRouter := mux.NewRouter().StrictSlash(true)

	testRouter.HandleFunc("/", homePage)
	testRouter.HandleFunc("/posts", allPosts).Methods("GET")
	testRouter.HandleFunc("/companies", allCompanies).Methods("GET")
	testRouter.HandleFunc("/user/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", testRouter))
}

//program starts from here
func main() {
	handleRequests()
}
