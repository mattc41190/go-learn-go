package main

import (
	"flag"
	"fmt"
)

// This program sets the value of flag regardless of its previous state. It is usable after parse.
func main() {
	favFoodPtr := flag.String("food", "No Answers Provided", "Your most favoritest food stuff")
	flag.Parse()
	fmt.Printf("My favorite food is: %s\n", *favFoodPtr)
	flag.Set("food", "Dots Candy!")
	fmt.Printf("NO! Shut up! Your favorite food is: %s\n", *favFoodPtr)
}
