package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	// Get a record of when the program started
	start := time.Now()

	// Create a channel for strings
	ch := make(chan string)

	// Loop over all CLI arguments
	for _, url := range os.Args[1:] {

		// Create a goroutine for the fetch function.
		// Passing it the channel we just made and our current CLI arg.
		go fetch(url, ch)

	}

	// Loop over all CLI arguments
	for range os.Args[1:] {

		// Wait for the channel to unload a value from the channel and then print that value
		fmt.Println(<-ch)

	}

	// Print the difference between when the program started and when it ended
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Decalred a function called fetch
// Which accepts a string and a channel that will have strings placed on to it
func fetch(url string, ch chan<- string) {

	// Get a record of when the function started
	start := time.Now()

	// If the string is not prefixed with "http://"
	if !strings.HasPrefix(url, "http://") {

		// add http as a prefix to the passed string.
		url = "http://" + url
	}

	// Ask the http module to get the content at the url we passed in.
	// Save the error value as well
	resp, err := http.Get(url)

	// If the error value is not `nil`
	if err != nil {

		// Put that error value on to the channel we passed in
		ch <- fmt.Sprint(err)

	}

	// Copy the contents from from the `Body` of the response to a void
	// Copy will return how many bytes it read. Save this value as `nbytes`
	// Copy will also return an error value. Overwrite our current err `var` with this value
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	// Since io.Copy opened the `Body` of the response to read it. We close it to avoid leaking resources.
	resp.Body.Close()

	// If the current error value is not `nil`
	if err != nil {

		// Put that error value (along with a small bit of context) on to the channel we passed in
		ch <- fmt.Sprintf("fetch reading: %v\n", err)

	}

	// Get the difference between when we started the request and what time it is now
	secs := time.Since(start).Seconds()

	// Create a string with the following data:
	// - The number of seconds for this request to finish.
	// - The number of bytes read from this request.
	// - The url we were getting this data for
	ch <- fmt.Sprintf("%.2fs %d %s", secs, nbytes, url)
}
