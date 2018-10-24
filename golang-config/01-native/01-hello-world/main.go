package main

import (
	"flag"
	"fmt"
)

func main() {

	// Declare variable set to the command line argument `--planet`.
	// It has a default value of Earth
	// If you run `./<bin> --help` you will see:
	// -planet string
	//   The planet you wish to greet (default "Earth")
	planetPtr := flag.String("planet", "Earth", "The planet you wish to greet")
	timesPtr := flag.Int("times", 1, "Number of times to greet the planet")

	// Parse must be passed for the values to be loaded in.
	flag.Parse()

	// Loop through a print line `*timesPtr` times
	for i := 0; i < *timesPtr; i++ {
		fmt.Println("Hello", *planetPtr)
	}
}
