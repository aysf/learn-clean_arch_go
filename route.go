package main

import (
	"encoding/json"
	"fmt"
	"gofb/entity"
	"gofb/repository"
	"math/rand"
	"net/http"
)

var (
	// posts []Post
	repo repository.PostRepository = repository.NewPostRepository()
)

// func init() {
// 	posts = []Post{
// 		{Id: 1, Title: "title 1", Text: "text 1"},
// 	}
// }

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	posts, err := repo.FindAll()
	fmt.Println("hanlder getPosts res:", posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "error getting posts"}`))
		return
	}

	// result, err := json.Marshal(posts)
	// if err != nil {
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	res.Write([]byte(`{"error": "error marshaling the post array"}`))
	// 	return
	// }

	res.WriteHeader(http.StatusOK)
	// res.Write(result)
	json.NewEncoder(res).Encode(posts)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var post entity.Post

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "error unmarshalling request"}`))
	}
	post.Id = rand.Int63()
	// posts = append(posts, post)
	repo.Save(&post)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)

	// result, _ := json.Marshal(posts)
	// res.Write(result)
}
