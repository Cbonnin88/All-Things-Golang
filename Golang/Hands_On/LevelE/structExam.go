package main

import "fmt"

type person struct {
	FirstName	string
	LastName	string
	IceCream	[]string
}
func main() { //nolint:typecheck
	me := person {
		FirstName: "Christopher",
		LastName: 	"Bonnin",
		IceCream: 	[]string{
					"chocolate",
					"peanut butter",
					"vanilla",
		},
	}

	julien := person {
		FirstName: "Julien",
		LastName: 	"Bonnin",
		IceCream: 	[]string {
					"vanilla",
					"chocolate",
					"mango",
		},
	}

	benjamin := person {
		FirstName: "Benjamin",
		LastName: "Calais",
		IceCream: []string{
					"vanilla",
					"pistache",
					"red fruit",
		},
	}

	// Creating a Map from our type 'person' struct
	lastName := map [string]person {
			me.LastName: me,
			julien.LastName: julien,
			benjamin.LastName: benjamin,
	}

	 firstName := map [string]person {
		julien.FirstName: julien,
		benjamin.FirstName: benjamin,
		me.FirstName: me,
	}
	// Ranging over the ice cream flavors
	fmt.Println("Ranging over our lastName map:")
	for _, v := range lastName {
		fmt.Println(v.FirstName)
		fmt.Println(v.LastName)
		for i, val := range v.IceCream {
			fmt.Println(i, val)
		}
		fmt.Println("")
	}

	fmt.Println("Raining over our firstNam map:")
	for _, v2 := range firstName {
		fmt.Println(v2.LastName)
		fmt.Println(v2.FirstName)
		for i, val2 := range v2.IceCream {
			fmt.Println(i,val2)
		}
		fmt.Println("")
	}

	fmt.Println(me.FirstName,me.LastName)
	for i, v := range me.IceCream {
		fmt.Println(i,v)
	}
	fmt.Println("")

	fmt.Println(julien.LastName,julien.FirstName)
	for i, v := range julien.IceCream {
		fmt.Println(i,v)
	}
	fmt.Println("")

	fmt.Println(benjamin.FirstName,benjamin.LastName)
	for i, v := range benjamin.IceCream {
		fmt.Println(i,v)
	}

}
