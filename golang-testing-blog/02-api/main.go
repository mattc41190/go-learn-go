package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Todo : A simple todo structure
type Todo struct {
	UserID    int    `json: userId`
	ID        int    `json: id`
	Title     string `json: title`
	Completed bool   `json: completed`
}

func makeItProud(s *string) {
	*s = strings.ToUpper(*s) + "!!!"
}

func getProudTodoTitle(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	t := Todo{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Fatal(err)
	}
	makeItProud(&t.Title)
	return t.Title
}

func main() {
	proudTitle := getProudTodoTitle("https://jsonplaceholder.typicode.com/todos/1")
	log.Printf(proudTitle)
}
