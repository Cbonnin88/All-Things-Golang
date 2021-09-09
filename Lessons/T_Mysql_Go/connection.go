package main

import (
	"database/sql"
	f "fmt" // referencing a package
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// I'm opening up a database connection using sql.Open
	// The database is called 'golangDB'
	// NOT the actual password for my database
	db, err := sql.Open("mysql", "root:pwd@tcp("+
		"localhost:3306)/university_records")

	//  Handle an error in case there's a problem with the connection
	if err != nil {
		panic(err.Error())
	}

	// Deferring the close until after the main function has finished
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	f.Println("Successfully connected to your Database !")
}
