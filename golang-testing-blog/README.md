# golang-testing-blog

## What is this?

This repo is a collection of insights gleaned from two blogs posts by (Alex Ellis) and (Nathan Leclaire) about testing programs in GoLang. Essentially it is a manually copied, revised, and (possibly) added to sample program written by the authors listed with tests included. 

## Why make this?

To begin exposing myself to how tests should be written in GoLang and to start getting familiar with how one might go about implementing TDD in GoLang. Some specific things I learned.

- Test tables are pretty great
- Not having native `assert` method is awkward
- Testing in GoLang is pretty straight forward.

## How do I run this?

Simply clone this repo `git clone <this-repo>` move into this directory `cd <path/to/example>` and run `go run main.go` to see the program implemented, or run `go test -v` to see the tests in action.  


## Resources: 

- https://blog.alexellis.io/golang-writing-unit-tests/

- https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/