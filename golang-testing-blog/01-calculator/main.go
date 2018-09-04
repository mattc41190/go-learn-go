package main

import (
	"log"
	"os"
	"strconv"
)

func getValues() (int, int) {
	a, err := strconv.Atoi(os.Args[1:][0])
	b, err := strconv.Atoi(os.Args[1:][1])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}

func main() {
	firstNum, secondNum := getValues()
	c := calculator{}
	num := c.Sum(firstNum, secondNum)
	log.Println(num)
}
