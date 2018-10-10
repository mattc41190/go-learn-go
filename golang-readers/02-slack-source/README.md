# 02-slack-source

## What is this?

This is code written by Andrei Tudor Călin and annotated by myself in hopes to see a fantastic means by which one may employ interfaces for generic programming in GoLang.

## Why make this?

To better understand the core concepts of GoLang and the Reader interface.

## Upshot -- Highly Personal, Stream of Consciousness 

- Making your data source's in `io.Reader`s helps you to be able to write a variety of sources. Consider the `io.Copy` method it already knows how to Read from an implementer of Read and place that data safely into anything implements the `io.Writer` interface. 

- The data will ALWAYS be held at the `io.Reader` level. Ergo, an implementation of Read can very safely assume that Read will work correctly if that convention is being followed. I think one of my main confusions was surrounding why I couldn't just work work with the data in question and the entire is POINT is a generic methodology of interacting with multiple sorts of _things that read_ (what I wanted to say was strings, but that isn't always how it works).

## Context

A Question to The Gopher Slack Channel:

Hey folks I am definitely a GoLang newbie so please be patient with my silliness, but I am having a very tough wrapping my head around GoLang interfaces. To be very specific I am having difficulty understanding the utility of the `io.Reader` interface. I understand what an `io.Reader` is, insomuch as it is an interface whose contract dictates a single `func` called `Read` be present which accepts a `slice` of bytes and returns an `int` and an `error`. 

What I do not understand is where this _methodology(?)_ becomes useful. What benefit do I gain from creating a `Reader`?  

What sorts of `Reader`s do you folks have running around in the wild and what would the alternative to these implementations have been?

Additionally, I have read the following resources with little more insight given to why `Reader`s are such a good idea, if you have any other sources or books related to this topic I would be very grateful. 

https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b
https://nathanleclaire.com/blog/2014/07/19/demystifying-golangs-io-dot-reader-and-io-dot-writer-interfaces/
https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185
https://play.golang.org/p/nxr38B5b_NC

Caveats: 

If this question is too simple, i.e. Along the lines of someone asking why 2 + 2 = 4, please tell me so and I will endeavor to take it on faith that `Reader`s == Good

If the question is less about `Reader`s and more about GoLang interfaces at large please let me know. 



Thanks in advance for any insight, 

Matt


----

@Matthew Cale Hello and welcome. Think of an `io.Reader` as a data source. It can be a socket, a file, or an in-memory buffer. An `io.Reader` is something you can read data from. The data is transferred to the `[]byte` you pass into the `Read` call.

Implementing your own `io.Reader` isn't very common, but making use of existing implementations definitely is.
By writing code that works on `io.Reader`, you are engaging in a form of generic programming. You're saying "my code can do something so long as it has access to some data stream".


That data stream can come from anywhere: a plain TCP socket, an HTTP response, a file on disk, or a buffer in memory.
Consider, for example `json.Decoder`. The decoder can decode JSON to Go types. All you need to do is feed it some data (wherever the data may come from).
It expresses this by saying it depends on an `io.Reader`.

https://play.golang.org/p/bux7OKwBP8h here's a simple implementation of `io.Reader`, which decorates another reader with additional behavior.
For a more involved and fundamental example, see: https://golang.org/pkg/bytes/#Reader. Better still, write code that uses readers.

To make this read [of the playground code] effective, check the documentation for the types and functions I use, such as `strings.Reader`, `io.Copy`, `ioutil.Discard`.

Andre

----

Andre,

The metaphor of an `io.Reader` as a data source is helpful. If in my mind I replace the word `io.Reader` with the word file for instance the concept gains clarity.

However the concept of engaging in generic programming still escapes me. When you say "my code can do something so long as it has access to some data stream" I lose the thread. In the case of `io.Reader` what is the "my code can do something", sorry if that is a hard read!

---

(After seeing golang playground code posted)

Andre,

The full read through and collection of my thoughts may take a bit will report back when done. Thanks for being responsive and patient