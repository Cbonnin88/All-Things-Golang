package main

import "fmt"

func main() {

	fmt.Println("On the", 1, "day of Christmas my true love sent to me:")
	fmt.Println("A partridge in a pear tree")
	fmt.Println("")

	for day := 2; day <= 12; day++ {
		fmt.Println("On the", day, "day of Christmas, my true love sent to me:")
		switch day {
		case 12:
			fmt.Println("Twelve Drummers drumming,")
			fallthrough
		case 11:
			fmt.Println("Eleven pipers piping,")
			fallthrough
		case 10:
			fmt.Println("Ten lords a-leaping,")
			fallthrough
		case 9:
			fmt.Println("Nine ladies dancing,")
			fallthrough
		case 8:
			fmt.Println("Eight maids a-milking,")
			fallthrough
		case 7:
			fmt.Println("Seven swans a-swimming,")
			fallthrough
		case 6:
			fmt.Println("Six geese a-laying,")
			fallthrough
		case 5:
			fmt.Println("Five golden rings,")
			fallthrough
		case 4:
			fmt.Println("Four calling birds,")
			fallthrough
		case 3:
			fmt.Println("Three french hens,")
			fallthrough
		case 2:
			fmt.Println("Two turtle doves, and")
			fallthrough
		case 1:
			fmt.Println("A partridge in a pear tree")

		}
		fmt.Println("")

	}
}
