package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Envelope struct {
	Kind string
	Msg  interface{}
}

type Sound struct {
	Description string
	Authority   string
}

func main() {

	// Create a place holder for raw json
	var msg json.RawMessage

	// The file variable is a pointer to an os.File
	// os.File is an implementer of Reader and thusly contains a Read function
	file, err := os.Open("data.json")

	if err != nil {
		panic(err)
	}

	// Read all the bytes out of the file into a variable called bytes
	input, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	// Create an Envelope with an empty value for Kind and a pointer to a json.RawMessage for Msg
	env := Envelope{
		Msg: &msg,
	}

	// Place the data from the read in JSON file (that was converted into bytes) into a known struct
	// Something important happens here...
	if err := json.Unmarshal(input, &env); err != nil {
		log.Fatal(err)
	}

	switch env.Kind {
	case "sound":
		var s Sound
		if err = json.Unmarshal(msg, &s); err != nil {
			panic(err)
		}
		desc := s.Description
		fmt.Println(desc)
	}

}

// Questions:
// What is a RawMessage?
