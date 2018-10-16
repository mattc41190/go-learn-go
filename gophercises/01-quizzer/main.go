package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type questionAndAnswer struct {
	question string
	answer   string
}

func main() {

	// Create flags for file name and time limit
	csvFileName := flag.String("csv", "problems.csv", "A CSV file with in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 10, "Time limit for the quiz in seconds")
	flag.Parse()

	// Read the file from disk
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open CSV: %s", *csvFileName)
		panic(err)
	}

	// Create quiz timer
	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))

	// Create a `csv` Reader from the File
	fileReader := csv.NewReader(file)

	// Create a populate a collection of Questions and Answers
	questionsAndAnswers := makeQuestionsAndAnswers(fileReader)

	// Ask the user each question and have them provide answers
	correctCount := 0
	answerReader := bufio.NewReader(os.Stdin)
	for _, entry := range questionsAndAnswers {
		answerChannel := make(chan bool)
		go func(entry questionAndAnswer, answerReader *bufio.Reader) {
			answerChannel <- askQuestion(entry, answerReader)
		}(entry, answerReader)
		select {
		// If there is a value to read off of the timer.C channel our time limit is up!
		case <-timer.C:
			fmt.Println("Bzzzz! Timer Expired!")
			provideSummary(float32(correctCount), float32(len(questionsAndAnswers)))
			return
		case isCorrect := <-answerChannel:
			if isCorrect {
				correctCount++
			}
		}
	}

	// Review the quiz with the user
	provideSummary(float32(correctCount), float32(len(questionsAndAnswers)))

}

func makeQuestionsAndAnswers(r *csv.Reader) []questionAndAnswer {
	questionsAndAnswers := []questionAndAnswer{}
	for {
		line, err := r.Read()
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

	return questionsAndAnswers
}

func askQuestion(entry questionAndAnswer, answerReader *bufio.Reader) bool {
	fmt.Printf("%v : ", entry.question)
	answer, err := answerReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	answer = strings.TrimSpace(answer)

	if answer == entry.answer {
		fmt.Println("Correct")
		return true
	}

	fmt.Println("Incorrect")
	fmt.Printf("Correct Answer: %v \n", entry.answer)
	fmt.Printf("You Answer: %v \n", answer)
	return false
}

func provideSummary(correctCount float32, total float32) {
	fmt.Println("----- Summary -----")
	// fmt.Println("Score:", (float32(correctCount)/float32(len(questionsAndAnswers)))*100)
	fmt.Println("Score:", (correctCount/total)*100)
	fmt.Println("Correct Count: ", correctCount)
	fmt.Println("Total Questions: ", total)
}
