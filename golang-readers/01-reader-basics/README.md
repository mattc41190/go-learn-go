# 01-reader-basics

## What is this?

This is the simplest most detailed version of a reader demo I could come up with. It will answer the following questions:

- What is a Reader?
    - A `Reader` from the `io` package is a GoLang interface which enforces the implementation of a single `func`. `Read`
    - https://golang.org/pkg/io/#Reader
- What does a `Reader.Read` look like and do?
    - A type which may call itself a `Reader` implements a `Read` function according to the following format:
        - `func (mt MyType) Read(b []byte) (int, error) {return <int>, <error>}`
    - The `Read` function, *generally* speaking, reads information into an `slice` of type `byte`.
        - The `Read` function will read from some source (stdin, file, HTTP response) and write that information to a byte slice.
        - This is what our implementation does in the `main.go` file of this project 

## Why make this?

More than anything the guides I have for understanding this are total shit. So here is my addition. -- poop emoji...

## How do I run this?

- `go run main.go`


