package main

import "fmt"

func GPA(finalGrade, midTermPaper float32) (string, float32) {
	var homeroomScore float32 = 17.5
	averageGrade := (finalGrade + midTermPaper + homeroomScore) / 2
	var letterGrade string
	switch {
	case averageGrade == 100:
		letterGrade = "A+"
	case averageGrade < 100 && averageGrade >= 90:
		letterGrade = "A"
	case averageGrade >= 80 && averageGrade < 90:
		letterGrade = "B"
	case averageGrade >= 70 && averageGrade < 80:
		letterGrade = "C"
	case averageGrade >= 65 && averageGrade < 70:
		letterGrade = "D"
	default:
		letterGrade = "F"
	}
	return letterGrade, averageGrade
}

func main() {
	var myMidTermPaper, myFinalGrade float32
	myMidTermPaper = 87.6
	myFinalGrade = 67.9
	var myAverage float32
	var myGrade string
	myGrade, myAverage = GPA(myMidTermPaper, myFinalGrade)
	fmt.Printf("I got a %v in History 213 and my average this year is %v ", myGrade, myAverage)
}
