// Original Source: Andrei Tudor Călin -- Minor modifications made
package main

import (
	"io"
	"os"
	"strings"
)

// We start by creating a type called countingReader
// countingReader has two fields src (which is an implementer of `io.Reader`)
// and total (which is an int64)
type countingReader struct {
	src   io.Reader
	total int64
}

// We declare a new function which accepts
// any implementer of `io.Reader` and returns a pointer
// to a value of type counterReader
func newCountingReader(r io.Reader) *countingReader {
	// We return a pointer to a literal countingReader.
	// `src` is set to the implementer of `io.Reader` that was passed in
	// and `total` is set to 0
	return &countingReader{
		// The line below (src:r) means that newCountingReader can implement
		// `io.Reader` simply by declaring the `Read` function (with the proper signature)
		// and passing the values through to the `src` attribute's `Read` function
		src:   r,
		total: 0,
	}
}

// We declare a function Read on a pointer to countingReader which accepts
// a byte slice argument and returns an int and an error.
// This function causes countingReader to be seen as an `io.Reader`.
func (cr *countingReader) Read(b []byte) (int, error) {
	// We declare vars `n` and `err` whose values will be set by invoking
	// the countingReader's src attribute's Read method on the byte slice passed.
	// Essentially, it is a pass through to a more "legitimate" Reader.
	n, err := cr.src.Read(b)
	// We set the countingReader's (cr's) `total` attribute to `n`
	// which will be the number of bytes read from the data source
	// provided to `cr.src.Read`
	cr.total += int64(n)
	// Adhereing to the `io.Reader` interface we return the number of bytes read
	// and any error that occurred on the underlying read.
	return n, err
}

func (cr *countingReader) Total() int64 {
	return cr.total
}

func main() {
	s := "hello" + "\n"
	// We create a new `string.Reader` (implementer of `io.Reader`) called `r`
	// by passing our string `s` to the strings package's function `NewReader`.
	// NOTE: `strings.NewReader()` is covered in detail in my article
	r := strings.NewReader(s)
	// We create a new countingReader called `cr`
	// by passing our reader `r` to our function `newCountingReader`.
	// Recall that countingReader's `src` attribute is set to the
	// implementer of `io.Reader` passed into `newCountingReader`
	cr := newCountingReader(r)
	// We invoke the io package's `Copy` function passing in `os.Stdout`
	// as the the `dst Writer` and our `countingReader` (`cr`) as the `src Reader`
	// NOTE: `ioutil.Copy` is covered in detail in my article
	// Previous Annotation: Consume all the data. -- Andrei's annotation
	io.Copy(os.Stdout, cr)
}
