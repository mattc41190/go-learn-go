package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

// Sometimes, if you are super weird, you will want a flags existence to make a difference
// Even if it is not set, just it's being present may shake things up. Consider the following.
func main() {

	// If I run this flag with as -h
	// The app will crash since no empty value, but present flag override exists
	hostPtr := flag.StringP("host", "h", "localhost", "The host you want to run this not web app on")
	// If I run this flag with as -p
	// There will be a change in the application and the understood default will NOT be used
	// In this case case the port's new default will be 7070
	portPtr := flag.StringP("port", "p", "8000", "The port you want to run this not web app on")
	flag.Lookup("port").NoOptDefVal = "7070"

	flag.Parse()

	fmt.Printf("Running on Host: %s Port: %s\n", *hostPtr, *portPtr)
}
