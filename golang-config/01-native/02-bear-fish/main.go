package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Args do not need to be declared prior to parsing, unlike flags.
	flag.Parse()
	// Get user input and set it to lowercase
	animal := strings.ToLower(flag.Arg(0))
	// Handle input
	switch animal {
	case "bear":
		fmt.Println("I am such a good bear. I want a tasty fish!")
	case "fish":
		fmt.Println("I am fish. Swim swim swim...")
	default:
		fmt.Printf("Program expects to be passed either fish or bear. You passed: %s\n", animal)
	}

}
