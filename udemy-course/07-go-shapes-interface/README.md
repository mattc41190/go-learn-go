# go-shapes-interface
### Creating and implementing interface in GoLang

## What is this?

`go-shapes-interface` is a tutorial application that, when run, will calculate the area of a square and triangle based on hardcoded parameters. 

## Why make this?

This application was made to complete a challenge in Stephen Grider's complete GoLang Complete Developer course on Udemy.

Much like the programs made for sections five and six, `go-shapes-interface` demonstrates GoLang's interface functionality. In this lesson you learn how to:

- Create an interface `type shape interface { getArea float64 }`
- Create types `type square struct { sideLength float64}`
- Create a func which implements an interface `func (s square) getArea() float64 {return ...}`

For more in depth explanation of interfaces see sections five and six. 

## How do I use this

The easiest way to use this application is simply by cloning it to your `$GOPATH/src` directory and running:
- `go run main.go`