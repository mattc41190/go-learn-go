package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "The URL path: %q\n", req.URL.Path)
}
