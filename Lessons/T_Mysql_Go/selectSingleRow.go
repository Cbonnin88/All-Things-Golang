package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Student struct {
	StudentID int
	FirstName string
	LastName  string
	Age       int
	Major     string
}

// Created my type Student variable
var studentInfo Student

func main() {
	db, err := sql.Open("mysql", "root:pwd@tcp("+
		"localhost:3306)/university_records")

	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Here I use an user input based on the age field to return my sql query
	var id int
	_, err = fmt.Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	// Querying a single row:
	err = db.QueryRow("SELECT FirstName, LastName, Age, Major FROM Students where StudentID = ?", id).Scan(&studentInfo.FirstName, &studentInfo.LastName,
		&studentInfo.Age, &studentInfo.Major)
	if err != nil {
		panic(err.Error())
	}

	log.Println(":", studentInfo.FirstName)
	log.Println(":", studentInfo.LastName)
	log.Println(":", studentInfo.Age)
	log.Println(":", studentInfo.Major)
}
