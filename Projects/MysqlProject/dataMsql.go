package main

import (
	"encoding/json"
	"log"
)

const info = `[{"student_ID": 1234, "first-name": "Alison", 
"last-name": "Fields", "age": 22, "major": "Biology"},
{"student_ID": 1235, "first-name": "Marcus", "last-name": "Jackson", 
"age": 18, "major": "Business Finance"},
{"student_ID": 1236, "first-name": "Rachel", "last-name": "Dowes", "age": 24,
"major": "Poltical Science"},
{"student_ID": 1237, "first-name": "Amos", "last-name":"Vanderkput", 
"age": 30, "major": "Culinary Arts"},
{"student_ID": 1238, "first-name": "Jenna", "last-name": "Wallace", 
"age": 40, "major": "Art"},
{"student_ID": 1239, "first-name": "Matthew", "last-name": "Baker", 
"age": 29, "major": "Gender Studies"},
{"student_ID": 1240, "first-name": "Laura", "last-name": "Croft", "age": 21, 
"major": "Geology"},
{"student_ID": 1241, "first-name": "Ethan", "last-name": "Winters", 
"age": 36, "major": "Neuroscience"},
{"student_ID": 1242, "first-name": "Alcina", "last-name": "Dimitrescu", 
"age": 44, "major": "Hospitality"}]`

func GetData () (students []Student) {
	err := json.Unmarshal([]byte(info), &students)
	if err != nil {
		log.Fatal(err)
	}
	return students
}
