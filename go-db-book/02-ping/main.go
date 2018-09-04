package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Registers itself with the sql package
)

var USER = "root"
var PASSWORD = os.Args[1:][0]
var ADDRESS = "tcp(localhost:3306)/test"
var CONNECTION_STRING = fmt.Sprintf("%s:%s@%s", USER, PASSWORD, ADDRESS)

func main() {

	db, err := sql.Open("mysql", CONNECTION_STRING)

	if err != nil {
		log.Fatal("Failed to connect")
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal("Ping failed")
	} else {
		log.Printf("Pong")
	}
}
