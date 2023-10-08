package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func init() {
// 	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/ananto/Sandbox/goplay/learn_go-crash-course-1/cred/fb.json")
// 	if err != nil {
// 		// Handle error if setting the environment variable fails.
// 		panic(err)
// 	}
// }

func main() {
	r := mux.NewRouter()
	const port = ":8000"
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "up and running")
	})

	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("server listening on port ", port)
	http.ListenAndServe(port, r)
}
