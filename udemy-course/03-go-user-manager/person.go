package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%v \n", p)
}

// NOTE: When type is *pointer and a struct/value -
// of that type is passed the pointer variable is created for you
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
