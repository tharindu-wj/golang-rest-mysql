package handler

import (
	"../db"
	"../model"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

//get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	result := db.FindAll("users")

	user := model.User{}
	users := []model.User{}

	for result.Next() {
		var id, company int
		var name, email, created_at string
		err := result.Scan(&id, &name, &email, &company, &created_at)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Name = name
		user.Email = email
		user.Company = company
		user.Created = created_at

		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)

}

//create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	items := map[string]string{
		"name":    r.Form.Get("name"),
		"email":   r.Form.Get("email"),
		"company": r.Form.Get("company"),
	}

	result := db.Save("users", items)

	if (result) {
		json.NewEncoder(w).Encode("New User Created")
	}
}

//get single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	result := db.FindBy("users", userID)

	user := model.User{}

	for result.Next() {
		var id, company int
		var name, email, created_at string
		err := result.Scan(&id, &name, &email, &company, &created_at)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Name = name
		user.Email = email
		user.Company = company
		user.Created = created_at
	}

	json.NewEncoder(w).Encode(user)
}
