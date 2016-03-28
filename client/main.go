package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	data := struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}{
		ID:      uuid.NewV4().String(),
		Title:   "",
		Content: "",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://127.0.0.1:4140/posts", "application/json", &buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Code: %d\n", resp.StatusCode)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
