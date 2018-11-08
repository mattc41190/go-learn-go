package main

import "fmt"

import flag "github.com/spf13/pflag"

// This program shows that generally speaking the native flags and pflags are interchangeable
func main() {

	// Most common way of making a CLI flag -- Get a pointer to string
	planetPtr := flag.String("planet", "Earth", "The planet you would like to greet")

	// Less common way of making CLI flag -- assigns value in flag to value variable
	var times int
	flag.IntVar(&times, "times", 1, "The number of times you would like ot greet the planet")

	// After all flags have been attached to flag we parse them
	flag.Parse()

	for i := 0; i < times; i++ {
		fmt.Printf("Hello %s\n", *planetPtr)
	}
}
