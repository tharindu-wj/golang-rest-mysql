package handler

import (
	"../db"
	"../model"
	"encoding/json"
	"net/http"
)

//get all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	db := db.DBConn()
	selDB, err := db.Query("SELECT * FROM posts ORDER BY created_at DESC")
	if err != nil {
		panic(err.Error())
	}
	post := model.Post{}
	posts := []model.Post{}
	for selDB.Next() {
		var p_id, u_id int
		var title, body, created_at string
		err = selDB.Scan(&p_id, &title, &body, &u_id, &created_at)
		if err != nil {
			panic(err.Error())
		}
		post.P_Id = p_id
		post.Title = title
		post.Body = body
		post.U_Id = u_id
		post.Created = created_at
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)

	defer db.Close()
}
