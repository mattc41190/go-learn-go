package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesWithDupes := []string{}
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Dup 2 Program could not open file -- Error %v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	if len(filesWithDupes) > 0 {
		fmt.Printf("Dupes found in: %s\n", filesWithDupes)
	}
}

func countLines(f *os.File, c map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		c[input.Text()]++
	}
}

// Bonus: Print all files where duplicates are found
// I would do this by replacing the simple counts map with a struct
// I would keep a map of words, then the value would be computed total of
// Number of times used per file and number of times used throughout all files
