# go-website-status-checker
### Using GoLang to knock on the front door of a few websites

## What is this?

`go-website-status-check` is a tutorial application that, when run, continually check the response sent from a collection of URLs 

## Why make this?

This application was made while following along with Stephen Grider's complete GoLang Complete Developer course on Udemy. The primary purpose of this exercise was to introduce the user to the concept of concurrency in GoLang. 

Some of the concepts covered here are:

- How to create channels and specify a channel's "type":
    - `c := make(chan string)` 
- How to specify a "go routine" and why someone would want to pass a channel into a routine:
    - `go someFunc(c)`
- How to supply a channel with a message:
    - `c <- "Hello"`
- How to receive message from a populated channel:
    - `greeting := <- c`
    - The command above is blocking, and can cause infinite waits.
- How to declare and use function literals
    - `func(s string) { fmt.Println(s) }("Hi")`

## How do I use this

The easiest way to use this application is simply by cloning it to your `$GOPATH/src` directory and running:
- `go run main.go`
## Other Resources:

[Scope Issue Ref One](http://oyvindsk.com/writing/common-golang-mistakes-1)
[Scope Issue Ref Two](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#closure_for_it_vars)