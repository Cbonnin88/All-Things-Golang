package main

import "fmt"

func main() { //nolint:typecheck
	var likes = map[string][]string {
		"Bonnin": {"Video games", "Poptarts", "Candy"},
		"Thiard": {"Money", "Art", "Mansions"},
		"Locatelli": {"Money", "Board Games", "Beer"},
		"Viollet-Jiménez": {"Cats", "Money", "Wine"},
	}

	fmt.Println(likes)

	// Adding a value:
	likes["O'Bready"] = []string{"Visual Arts","Magic","cats"}
	delete(likes,"Bonnin")

	for key, val := range likes{
		fmt.Println("this is the record for:",key)
		for i, val2 := range val {
			fmt.Println("\t",i,val2)

		}
	}
}
