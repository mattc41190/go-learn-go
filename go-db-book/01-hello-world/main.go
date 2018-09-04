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
		log.Fatal(err)
	}
	defer db.Close()

	createTable := "CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))"

	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal(err)
	}

	insertData := "INSERT INTO test.hello(world) VALUES ('hello world')"

	res, err := db.Exec(insertData)

	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("inserted %d rows", rowCount)

	selectData := "SELECT * FROM test.hello"

	rows, err := db.Query(selectData)

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var s string

		err = rows.Scan(&s)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Found row containing %q", s)

	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
