package main

import (
	"fmt"
)

func main() {
	num := 5
	pointerToNum := &num
	fmt.Println(pointerToNum)  // A memory address
	fmt.Println(*pointerToNum) // The value held at that memory address

	*pointerToNum = 10 // A "read through" from the pointer to its value
	fmt.Println(num)   // 10!
}
