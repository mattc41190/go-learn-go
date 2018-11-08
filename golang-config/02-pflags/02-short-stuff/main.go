package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

// This program show the primary difference between Flags and PFlags
// Namely, there is a shorthand option in PFlags.
// Note, that StringP() accepts 4 args while String() accepts only 3.
// The fourth is the shorthand arg notation `go run main.go -t=Man || --title=Man`
func main() {
	namePtr := flag.StringP("title", "t", "Man", "Title, make it rhyme with Stan?")
	flag.Parse()
	fmt.Printf("Stan the %s\n", *namePtr)
}
