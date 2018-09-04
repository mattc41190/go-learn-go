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
		log.Fatal("Failed to connect", err)
	}

	tx, err := db.Begin()

	if err != nil {
		log.Fatal("Failed to create transaction", err)
	}

	updateStmt := "UPDATE test.hello SET world = ? WHERE world = ?"
	tx.Exec(updateStmt, "hello saturn", "hello mars")

	err = tx.Commit()

	if err != nil {
		log.Fatal("Failed to commit transaction", err)
	}
}
