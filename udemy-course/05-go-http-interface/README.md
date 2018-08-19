# go-http-interface
### First steps on using GoLang to read data from the web 


## What is this?

`go-http-interface` is a tutorial application that demonstrates the usage of interfaces in GoLang by creating an interface which will adhere to the interface laid out by the `io.Copy` function. 

## Why make this?

This application was made to complete a challenge in Stephen Grider's complete GoLang Complete Developer course on Udemy.

Its primary purpose is to demonstrate how to write code that will agree with the contract set out by certain standard library methods. In this case we are attempting to place data that we have received from the internet into stdout. To accomplish we are doing the following:

_Please note, this explanation contains quite a lot of detail_

- Importing and using `http.Get` to get data from the web.
    - [http.Get Docs](https://golang.org/pkg/net/http/#Client.Get)
- Saving the response from that call to a var called `resp`
- Once we have the `resp` we can inspect it and determine it has a field called `Body` that has implementa the `io.ReadCloser` interface.
    - [`Response`](https://golang.org/pkg/net/http/#Response) -- Contains Body struct which adheres to io.ReadCloser
    - [`io.ReadCloser`](https://golang.org/pkg/io/#ReadCloser)
        - [`Reader`](https://golang.org/pkg/io/#Reader)
        - [`Closer`](https://golang.org/pkg/io/#Closer)
    - Though the actual implementation of the Read method specified by the ReadCloser interface is a bit obscured the documentation for every implementation is suggested to be reading bytes into a byte slice (thereby manipuating the passed in underlying value) and to return an int (representing the number of bytes read) and an err struct 
- How to create a function which adheres to the [Writer](https://github.com/golang/go/blob/master/src/io/io.go#L90-L92) interface by creating a type called `logWriter`
    - `type logWriter struct{}`
    - `func (lw logWriter) Write(bs slice[]) (int, err) {}`
- How to use `io.Copy` to copy data from a Reader type to a Writer type.
    - In this case specifically the task was to create a type which would serve a replacement for os.Stdout

## How do I use this

The easiest way to use this application is simply by cloning it to your `$GOPATH/src` directory and running:
- `go run main.go`