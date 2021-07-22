package main

import "fmt"

// Our original struct type 'person':
type person struct {
	firstName string
	lastName  string
	age       int
}

// Here we are embedding our type 'person' into type 'politician' since a 'person' can also be a 'politician':
type politician struct {
	person
	polParty      string
	position      string
	yearsInOffice int
}

func main() { //nolint:typecheck
	p1 := politician{
		person: person{
			firstName: "Emmanuel",
			lastName:  "Macron",
			age:       43,
		},
		polParty:      "La Republique En Marche",
		position:      "President of France",
		yearsInOffice: 4,
	}

	p2 := politician{
		person: person{
			firstName: "Joe",
			lastName:  "Biden",
			age:       78,
		},
		polParty:      "Democratic Party",
		position:      "President of the United States",
		yearsInOffice: 1,
	}

	p3 := politician{
		person: person{
			firstName: "Angela",
			lastName:  "Merkel",
			age:       66,
		},
		polParty:      "Christian Democratic Union",
		position:      "Chancellor of Germany",
		yearsInOffice: 16,
	}

	p4 := politician{
		person: person{
			firstName: "Stefan",
			lastName:  "Lofven",
			age:       63,
		},
		polParty:      "Social Democrats",
		position:      "Prime Minister of Sweden",
		yearsInOffice: 7,
	}

	fmt.Println("Our embedded Structs:")

	fmt.Println("The leader of France is: ", p1)
	fmt.Println("The leader of the United States is: ", p2)
	fmt.Println("The leader of Germany is: ", p3)
	fmt.Println("The leader of Sweden is: ", p4, "\n")

	// We can also access fields within our embedded Structs:
	fmt.Println("Accessing fields within an embedded Struct:")
	fmt.Println("The first name of the President of France is:", p1.person.firstName)
	fmt.Println("President Biden current age is:", p2.person.age)
	fmt.Println("The Chancellor of Germany's last name is:", p3.person.lastName)
	fmt.Println("The Prime minister of Sweden's full name is:", p4.person.firstName, p4.person.lastName)

}
