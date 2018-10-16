package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))

}

/*

Bonus 01:

	fmt.Println(strings.Join(os.Args[:], " "))
	// /var/folders/yf/lj25jv3s26qb9vl4fxj2glkjstf48p/T/go-build489417870/b001/exe/main you are one cool guy

Bonus 02:

	var s, sep string
	for i, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(i) + arg
		sep = " "
	}
	fmt.Println(s)

Bonus 03:

	!!! TODO !!!
	!!! Chapter 01 -- Page 8 !!!

*/
