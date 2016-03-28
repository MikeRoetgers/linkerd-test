package main

import (
	"log"
	"net/http"
)

func main() {
	s := newService()
	http.HandleFunc("/", s.handleBlogPosts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
