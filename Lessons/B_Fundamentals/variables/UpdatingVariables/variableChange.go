package main

import "fmt"

func main() {

	var catsForSale float64 // catsForSale has not assigned value, it is understood to be "0.0"
	catPrice := 12.66

	catsForSale += catPrice // we add the two variables together, which will change the value of catsForSale

	fmt.Println(catsForSale) // will print "12.66" because catsForSale has been assigned the value of 12.66

	/* Ex:
	catsForSale's original (float) valule is 0.0 and catPrice is 12.66
	catsForSale = catsForSale (0.0) + catPrice(12.66)
	Go will print 12.66 because 12.66 + 0.0 is 12.66
	*/
	dogPrice := 20.12 // we add a new item to change once again the value of catsForSale
	catsForSale = dogPrice + catPrice
	fmt.Println(catsForSale)
	/* This time catsForSale will print 32.78, because dogPrice(20.12) + catPrice(12.66) equal 32.78
	 */
}

// variables are different from constants because we can update them
