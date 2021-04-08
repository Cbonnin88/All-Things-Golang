package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	file := "The Raven.txt"

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file is not found", file)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)

	pdf.MultiCell(190, 5, string(content), "0", "0", false)
	_ = pdf.OutputFileAndClose("The Revan.pdf")

	fmt.Println("Your poem was transformed into a PDF")

}
