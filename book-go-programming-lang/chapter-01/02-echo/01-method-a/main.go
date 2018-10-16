package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seperator string
	for i := 1; i < len(os.Args); i++ {
		// Clever trick using the init value of string and then resetting to a space each time after.
		// Not efficient, clever, but terse.
		s += seperator + os.Args[i]
		seperator = " "
	}
	fmt.Println(s)
}
