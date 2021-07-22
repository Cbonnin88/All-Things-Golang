package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	CompanyName       string  `json:"companyName"` // here we add JSON identifiers to our struct
	NumberOfEmployees int     `json:"number_of_employees"`
	Employee		  Employee `json:"employee"` // Our nested 'Employee' struct
	Country           string  `json:"country"`
	IsActive          bool    `json:"is_active"`
	Location          string  `json:"location"`
	FemaleManagers    float64 `json:"female_managers"`
}
type Employee struct { // here we added our json keys to our nested Employee struct
	Name	string 		`json:"name"`
	Age		int			`json:"age"`
	DOB     string      `json:"dob"`
	Title	string		`json:"title"`
	Department  string	`json:"department"`
	StartDate	 string `json:"start_date"`
	IsEmployed	  bool  `json:"is_employed"`

}
func main() {
	fmt.Println("our company struct:")
	// here we declare a new employee:
	employee := Employee{Name: "Janice Reed", Age: 32, DOB: "03/10/1988", Title: "Golang Developer", Department: "IT Department", StartDate: "10/07/2020", IsEmployed: true}

	company := Company{CompanyName: "Azfalte", NumberOfEmployees: 10, Employee: employee, Country: "France", IsActive: true, Location: "Lille", FemaleManagers: 50.0}
	fmt.Printf("%+v\n",company)
	fmt.Println("")

	byteArray, err := json.MarshalIndent(company,"","    ") // here we convert our Go object into a JSON string using the 'Marshal' function
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Our company struct converted to a JSON string:")
	fmt.Println(string(byteArray))
}

