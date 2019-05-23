package handler

import (
	"../db"
	"../model"
	"encoding/json"
	"net/http"
)

//get all companies
func GetCompanies(w http.ResponseWriter, r *http.Request) {
	result := db.FindAll("companies")
	company := model.Company{}
	companies := []model.Company{}
	for result.Next() {
		var id int
		var name, location, created_at string
		err := result.Scan(&id, &name, &location, &created_at)
		if err != nil {
			panic(err.Error())
		}
		company.Id = id
		company.Name = name
		company.Location = location
		company.Created = created_at
		companies = append(companies, company)
	}

	json.NewEncoder(w).Encode(companies)
}
