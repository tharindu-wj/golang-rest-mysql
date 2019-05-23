package handler

import (
	"../db"
	"../model"
	"encoding/json"
	"net/http"
)

//get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := db.DBConn()
	selDB, err := db.Query("SELECT * FROM users ORDER BY created_at DESC")
	if err != nil {
		panic(err.Error())
	}

	user := model.User{}
	users := []model.User{}

	for selDB.Next() {
		var u_id, c_id int
		var name, email, created_at string
		err = selDB.Scan(&u_id, &name, &email, &c_id, &created_at)
		if err != nil {
			panic(err.Error())
		}
		user.U_Id = u_id
		user.Name = name
		user.Email = email
		user.C_Id = c_id
		user.Created = created_at

		users = append(users, user)
	}
	defer db.Close()
	json.NewEncoder(w).Encode(users)

}
