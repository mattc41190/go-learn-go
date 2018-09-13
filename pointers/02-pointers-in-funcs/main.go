package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "matt"
	fixMyName(&name)
	fmt.Println("My name is", name)
}

func fixMyName(name *string) {
	*name = strings.Title(*name)
}
