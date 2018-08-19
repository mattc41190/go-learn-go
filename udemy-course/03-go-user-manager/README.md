# go-user-manager
# The very beginning of modeling user like structs in GoLang

## What is this?

`go-user-manager` is a CLI application that display the first steps in understanding GoLang's `struct` data type and provides a glimpse into accessing .

## Why make this?

This application was made while following along with Stephen Grider's complete GoLang Complete Developer course on Udemy. 

The primary purpose was introduce a new GoLang developer to the concept of how `struct`s work and to briefly introduce the concept of pointers in GoLang. Some of the concepts covered were:

- How to declare variables of type `struct`. 
    - `type Person struct {prop propType}`
- How to nest `struct`s.
- How to find the memory address that a variable "points" to in GoLang
    - `num := 5`
    - `pointerToNum := &num`
- How to modify the value stored at a particular memory address via pointers 
    - `num := 5`
    - `pointerToNum := &num`
    - `*pointerToNum = 10`
    - `num == 10 // true`
- How to use the GoLang shortcuts for pointers as receivers. See code or review this [video](https://www.udemy.com/go-the-complete-developers-guide/learn/v4/t/lecture/7797352?start=0).



## How do I use this

The easiest way to use this application is simply by cloning it to your `$GOPATH/src` directory and running:
- `go run main.go person.go`