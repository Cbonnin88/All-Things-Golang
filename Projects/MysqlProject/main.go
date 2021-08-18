package main

import (
	"database/sql"
	"fmt"
	"log"
)


const (
	db_user = "root"
	db_pwd = "Jujuetchris1989!"
	db_addr = "127.0.0.1"
	db_name = "GolangSql"
)

type Student struct {
	StudentID    int
	Firstname	 string
	Lastname	 string
	Age 		 int
	Major		 string
}

func main() {
	s := fmt.Sprintf("%s:%s@tcp(%s:3306/%s",
		db_user,db_pwd,db_addr,db_name)
	fmt.Println(s)
	db, err := sql.Open("mysql", s)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

}
