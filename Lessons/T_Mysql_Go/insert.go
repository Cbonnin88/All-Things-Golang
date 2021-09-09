package main

import (
	"database/sql"
	f "fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "root:pwd@tcp("+
		"localhost:3306)/university_records")

	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	// Performing a db.Query insert:
	insert, err := db.Query("INSERT INTO Students VALUES (19, 'Jake', 'Kelling', 50, 'Welding'),(20, 'Samantha', 'Thomas', 29, 'Carpentry');")

	// if there's an error with the insertion , handle it
	if err != nil {
		panic(err.Error())
	}

	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {
			f.Print("There was an error with inserting the data")

		}
	}(insert)

	f.Println("insert was successful")
}
