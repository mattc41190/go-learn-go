package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Oh no we got err")
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
