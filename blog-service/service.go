package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

type service struct {
	posts map[string]*BlogPost
}

func newService() *service {
	return &service{
		posts: make(map[string]*BlogPost),
	}
}

// BlogPost ...
type BlogPost struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (s *service) handleBlogPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	switch r.Method {
	case "POST":
		s.handleCreatePost(w, r)
	case "GET":
		s.handleGetPosts(w, r)
	default:
		http.Error(w, "Method not supported", 405)
		return
	}
}

func (s *service) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var post BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode JSON: %s", err), 400)
		return
	}
	id := uuid.NewV4().String()
	post.ID = id
	s.posts[id] = &post
	w.WriteHeader(http.StatusCreated)
}

func (s *service) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts := []*BlogPost{}
	for _, p := range s.posts {
		posts = append(posts, p)
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(posts)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode JSON: %s", err), 400)
		return
	}
	w.Write(buf.Bytes())
}
