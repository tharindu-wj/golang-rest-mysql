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
		var id, user int
		var title, body, created_at string
		err := result.Scan(&id, &title, &body, &user, &created_at)
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

}

//get single post data
func GetPost(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["id"]

	result := db.FindBy("posts", postID)

	post := model.Post{}

	for result.Next() {
		var id, user int
		var title, body, created_at string
		err := result.Scan(&id, &title, &body, &user, &created_at)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Title = title
		post.Body = body
		post.User = user
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
		"user" : r.Form.Get("user"),
	}

	result := db.Save("posts", items)

	if(result){
		json.NewEncoder(w).Encode("New Post Created")
	}
}

//save post
func RemovePost(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["id"]

	result := db.Remove("posts", postID)

	if(result){
		json.NewEncoder(w).Encode("Post Deleted")
	}
}