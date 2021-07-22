package main

import "fmt"

func main() { //nolint:typecheck
	bond := []string{"James","Bond","Shaken","not stirred"}
	moneyPenny :=[]string{"Miss","Moneypenny","Hellooo, James."}
	fmt.Println("James Bond slice:",bond)
	fmt.Println("Miss Moneypenny slice:",moneyPenny)

	mixed := [][]string{bond,moneyPenny}
	fmt.Println("Bond and Moneypenny:",mixed)

	for i, xs := range mixed {
		fmt.Println("record", i) // i represents our first slice
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t %v \n",j, val) // j reprends our second slice
		}
	}


}
