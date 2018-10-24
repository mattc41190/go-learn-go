package main

import (
	"flag"
	"fmt"
)

func main() {
	relativeAuntPtr := flag.String("aunt", "Becky", "Your aunt's name")
	relativeMomPtr := flag.String("mom", "Sally", "Your mom's name")
	relativeDadPtr := flag.String("dad", "Roland", "Your dad's name")

	fmt.Println("Visit all variables prior to actually parsing the commands into memory from os.Args")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("You %s has name of: %s\n", f.Name, f.Value)
	})

	flag.Parse()

	fmt.Println("Visit all variables after parsing the commands into memory from os.Args")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("You %s has name of: %s\n", f.Name, f.Value)
	})

	// Keeping this in to show that the typical flag usage is not modified based on above activity.
	// Feel free to ignore the code below.
	relatives := []*string{
		relativeAuntPtr,
		relativeDadPtr,
		relativeMomPtr,
	}

	for range relatives {
		fmt.Print()
	}

}
