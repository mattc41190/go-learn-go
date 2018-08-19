# go-bot-interface
### A very gentle introduction to interfaces in GoLang

## What is this?

`go-bot-interface` is a tutorial application that demonstrates the usage of interfaces in GoLang by creating fake "bots" that accomplish similar functionality but speak different languages. 

## Why makes this?

This application was made along side the same application in Stephen Grider's complete GoLang Complete Developer course on Udemy. 

It's primary purpose is to quickly expose a new GoLang developer to what interfaces are in GoLang, and how to use them. Some of the language concepts used in this example are:

- How to declare an interface in GoLang
    - `type interfaceName interface { functionName() [returnType]}`
- How to declare a function that implements an interface type
    - `func funcName(i interfaceName) [returnType] {i.functionName()}`

## How do I use this

The easiest way to use this application is simply by cloning it to your `$GOPATH/src` directory and running:
- `go run main.go`


