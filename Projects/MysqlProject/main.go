package main

import (
	"database/sql"
	f "fmt" // referencing a package
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// I'm opening up a database connection using sql.Open
	// The database is called 'golangDB'
	db, err := sql.Open("mysql", "root:Jujuetchris1989!@tcp("+
		"localhost:3306)/university_records")
	if err != nil {
		f.Println("Encountered error validating sql.Open arguments")
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	// Here I am verifying if the connection is still available
	err = db.Ping()
	if err != nil {
		f.Println("Unable to verify connection")
		panic(err.Error())
	}

	// Here I am inserting data into my university records:
	insert, err := db.Query("INSERT INTO Students(StudentID, FirstName, LastName, Age, Major) VALUES(10,'Alcina', 'Dimitrescu',44, 'Biology');")
	if err != nil {
		panic(err.Error())
	}
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(insert)
	f.Println("Successfully connected to your Database !")

}
