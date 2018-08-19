package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	larry := englishBot{}
	miguel := spanishBot{}
	printGreeting(larry)
	printGreeting(miguel)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (b englishBot) getGreeting() string {
	return "Hi There!"
}

func (b spanishBot) getGreeting() string {
	return "Hola!"
}
