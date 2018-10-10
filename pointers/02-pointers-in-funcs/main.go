package main

import (
	"fmt"
	"strings"
)

// Declare a function which accepts a value of type pointer to string
func fixMyName(name *string) {
	// Perform a "read-through" on the pointer to the value therein and modify it
	*name = strings.Title(*name)
}

func main() {
	name := "matt"
	fixMyName(&name)
	fmt.Println("My name is", name)
}
