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

	result := db.FindBy("posts", postID)

	post := model.Post{}

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
	}

	json.NewEncoder(w).Encode(post)
}

//save post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	items := map[string]string{
		"title" : r.Form.Get("title"),
		"body" :r.Form.Get("body"),
		"u_id" : r.Form.Get("u_id"),
	}

	result := db.Save("posts", items)

	if(result){
		json.NewEncoder(w).Encode("New Post Created")
	}
}