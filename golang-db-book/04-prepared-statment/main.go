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

	// NOTE: Prepared statements can cause memory and perf issues if not used carefully

	db, err := sql.Open("mysql", CONNECTION_STRING)

	if err != nil {
		log.Fatal("Failed to connect", err)
	}

	// Template Statement

	insertInfo := "INSERT INTO test.hello(world) VALUES(?)"

	// Example 1: A single "template" version of a SQL statement

	_, err = db.Exec(insertInfo, "hello mercury")

	if err != nil {
		log.Fatal("Failed single prepared statement", err)
	}

	// Example 2: A true "prepared" SQL statement
	// See gotchas and performance details on page 26 of:
	// https://itjumpstart.files.wordpress.com/2015/03/database-driven-apps-with-go.pdf

	stmt, err := db.Prepare(insertInfo)

	defer stmt.Close()

	if err != nil {
		log.Fatal("Failed to prepare stmt", err)
	}

	otherPlanets := []string{"hello venus", "hello earth", "hello mars", "hello jupiter"}

	for _, p := range otherPlanets {
		_, err := stmt.Exec(p)

		if err != nil {
			log.Fatal("Failed to insert greeting to planet", err)
		}
	}

}
