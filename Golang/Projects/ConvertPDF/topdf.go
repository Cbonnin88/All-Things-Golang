package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	file := "/Users/chris/Go/Golang/Projects/ConvertPDF/Orea Cheesecake Recipe"

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file is not found", file)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)
	_ = pdf.OutputFileAndClose("Orea Cheesecake.pdf")

	fmt.Println("Your recipe was transformed into a PDF")
}
