package handler

import (
	"../db"
	"../model"
	"encoding/json"
	"net/http"
)

//get all companies
func GetCompanies(w http.ResponseWriter, r *http.Request) {
	db := db.DBConn()
	selDB, err := db.Query("SELECT * FROM companies")
	if err != nil {
		panic(err.Error())
	}
	company := model.Company{}
	companies := []model.Company{}
	for selDB.Next() {
		var c_id int
		var name, location string
		err = selDB.Scan(&c_id, &name, &location)
		if err != nil {
			panic(err.Error())
		}
		company.C_Id = c_id
		company.Name = name
		company.Location = location
		companies = append(companies, company)
	}
	defer db.Close()
	json.NewEncoder(w).Encode(companies)
}
