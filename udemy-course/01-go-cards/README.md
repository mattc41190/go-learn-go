# go-cards
### A CLI application that behaves sort of like a deck of cards

## What is this?

`go-cards` is a tutorial application written in GoLang that simulates certain behaviors you would expect to find in a deck of cards. 

## Why make this?

This application was made along side the same application in Stephen Grider's complete GoLang Complete Developer course on Udemy. 

It's primary purpose is to introduce a new developer to several language constructs in GoLang. Some of the concepts that were introduced in this example are:

- How to use the the `go run <file>.go` tool
- How to use the the `go build <file>.go` tool
- How to set up the entry point to a GoLang application
    - `package main`
    - `func main () {}`
- How variables can be initialized in GoLang:
    - `var greeting = "hello"`
    - `farewell := "bye"`
- How functions can be declared in GoLang:
    - `func [(r receiverType)] funcName (a argType) [(returnType)] {[body]}`
- How to declare a new `type` in GoLang
- How to import items from the Standard Library
- How to declare a range based `for` loop
- How to work with the `slice` type.
    - Analogous to JavaScript `Array`
- How to use the `go fmt <file>.go` tool.
- How to return multiple values from a function.
- How to create Date based random numbers.
- How to call functions on a variable whose receiver is correct.
- How to take advantage of GoLang's built-in testing framework

## How do I use this

The easiest way to use this application is simply by cloning it to your `$GOPATH/src` directory and running:
- `go run main.go deck.go`
