package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	out, err := os.Create("index.html")
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal("file not found")
		}
	}(out)
	check(err)

	data := map[string]interface{}{}
	file, err := ioutil.ReadFile("Projects/resumeProject/data.json")
	check(err)

	err = json.Unmarshal(file, &data)
	check(err)
	t, err := template.ParseGlob("Projects/resumeProject/templateHTML/*")
	check(err)

	err = t.Execute(out, data)
	check(err)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
