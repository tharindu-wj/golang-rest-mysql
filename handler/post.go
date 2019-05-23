package handler

import (
	"../db"
	"../model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//get all posts w http.ResponseWriter, r *http.Request
func GetPosts(w http.ResponseWriter, r *http.Request) {
	result := db.FindAll("posts")
	post := model.Post{}
	posts := []model.Post{}
	for result.Next() {
		var p_id, u_id int
		var title, body, created_at string
		err := result.Scan(&p_id, &title, &body, &u_id, &created_at)
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

}

//get single post data
func GetPost(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["id"]
	db := db.DBConn()
	selDB, err := db.Query("SELECT * FROM posts WHERE p_id=?", postID)
	if err != nil {
		panic(err.Error())
	}

	post := model.Post{}

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
	}
	defer db.Close()
	json.NewEncoder(w).Encode(post)
}
