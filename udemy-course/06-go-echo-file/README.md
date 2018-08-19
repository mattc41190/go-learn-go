# go-echo-file
### First steps on using GoLang to read data from disk

## What is this?

`go-echo-file` is a tutorial application that demonstrates how to pass command line arguments to an application.

## Why make this?

This application was made to complete a challenge in Stephen Grider's complete GoLang Complete Developer course on Udemy. It builds almost entirely on section five's discussion on the http interface and how to write methods that adhere to and take advantage of interface implementations. One large difference is that in the http lesson the user gathers data from a hardcoded website whereas in this lesson the user pass in a command line argument specifying the (relative or absolute) path to a file.

For more details on the usage of interface see section five README.md.

## How do I use this

The easiest way to use this application is simply by cloning it to your `$GOPATH/src` directory and running:
- `go run main.go file.txt`