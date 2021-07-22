package main

import "fmt"

// A Struct is a data structure which allows us to compose values of different types:

type country struct {
	country      string
	capital      string
	officialLang string
}

// We can also have an 'empty' Struct
type empty struct{}

func main() { //nolint:typecheck
	France := country{
		country:      "France",
		capital:      "Paris",
		officialLang: "French",
	}

	Germany := country{
		country:      "Germany",
		capital:      "Berlin",
		officialLang: "German",
	}
	Finland := country{
		country:      "Finland",
		capital:      "Helsinki",
		officialLang: "Finnish",
	}

	UnitedKingdom := country{
		country:      "United Kingdom",
		capital:      "London",
		officialLang: "English",
	}
	fmt.Println("Our first struct is: ", France)
	fmt.Println("our second struct is: ", Germany)
	fmt.Println("our third struct is: ", Finland)
	fmt.Println("Our fourth struct is: ", UnitedKingdom)
	fmt.Println("An empty struct: ", empty{}, "\n")

	// You can access the fields of the Struct with .field name:
	fmt.Println("Accessing fields from our structs:")
	fmt.Println("The official language of France : ", France.officialLang)
	fmt.Println("The capital of Germany: ", Germany.capital)
	fmt.Println("The only nordic country that uses the Euro: ", Finland.country)
	fmt.Println("The language that isn't the official language of the United Kingdom: ", UnitedKingdom.officialLang)

}
