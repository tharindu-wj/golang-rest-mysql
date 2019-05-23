package handler

import (
	"../db"
	"../model"
	"encoding/json"
	"net/http"
)

//get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	result := db.FindAll("users")

	user := model.User{}
	users := []model.User{}

	for result.Next() {
		var u_id, c_id int
		var name, email, created_at string
		err := result.Scan(&u_id, &name, &email, &c_id, &created_at)
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

	json.NewEncoder(w).Encode(users)

}

//create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	items := map[string]string{
		"name" : r.Form.Get("name"),
		"email" :r.Form.Get("email"),
		"c_id" : r.Form.Get("c_id"),
	}

	result := db.Save("users", items)

	if(result){
		json.NewEncoder(w).Encode("New User Created")
	}
}
