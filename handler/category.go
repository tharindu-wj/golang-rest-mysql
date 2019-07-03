package handler

import (
	"../db"
	"../model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//get all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	result := db.FindAll("categories")
	category := model.Category{}
	companies := []model.Category{}
	for result.Next() {
		var id int
		var name, description, created_at string
		err := result.Scan(&id, &name, &description, &created_at)
		if err != nil {
			panic(err.Error())
		}
		category.Id = id
		category.Name = name
		category.Description = description
		category.Created = created_at
		companies = append(companies, category)
	}

	json.NewEncoder(w).Encode(companies)
}

//get single post
func GetCategory(w http.ResponseWriter, r *http.Request) {
	catID := mux.Vars(r)["id"]

	result := db.FindBy("categories", catID)

	category := model.Category{}

	for result.Next() {
		var id int
		var name, description, created_at string
		err := result.Scan(&id, &name, &description, &created_at)
		if err != nil {
			panic(err.Error())
		}
		category.Id = id
		category.Name = name
		category.Description = description
		category.Created = created_at
	}

	json.NewEncoder(w).Encode(category)
}
