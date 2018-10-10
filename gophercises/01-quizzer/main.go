package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type questionAndAnswer struct {
	question string
	answer   string
}

func main() {
	// Read the file from disk
	file, err := os.Open("./problems.csv")
	if err != nil {
		panic(err)
	}

	// Create a `csv` Reader from the File
	fileReader := csv.NewReader(file)

	// Create a populate a collection of Questions and Answers
	questionsAndAnswers := []questionAndAnswer{}
	for {
		line, err := fileReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		entry := questionAndAnswer{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
		questionsAndAnswers = append(questionsAndAnswers, entry)
	}

	correctCount := 0

	// Ask the user each question and have them provide answers
	answerReader := bufio.NewReader(os.Stdin)
	for _, entry := range questionsAndAnswers {
		fmt.Printf("%v : ", entry.question)
		answer, err := answerReader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		answer = strings.TrimSpace(answer)

		if answer == entry.answer {
			fmt.Println("Correct")
			correctCount++
		} else {
			fmt.Println("Incorrect")
			fmt.Printf("Correct Answer: %v \n", entry.answer)
			fmt.Printf("You Answer: %v \n", answer)
		}

	}

	fmt.Println("----- Summary -----")
	fmt.Println("Score:", (float32(correctCount)/float32(len(questionsAndAnswers)))*100)
	fmt.Println("Correct Count: ", correctCount)
	fmt.Println("Total Questions: ", len(questionsAndAnswers))

}
