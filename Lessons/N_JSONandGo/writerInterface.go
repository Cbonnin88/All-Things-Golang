package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Vehicle struct {
	Year	int
	Model   string
	IsManuel	bool
	Price		float64
	Tva 		float64
}

func main() {
	car := Vehicle{
			Year: 2021,
			Model: "Volkswagen",
			IsManuel: false,
			Price:	21000,
			Tva: 20.0,
	}

	// For testing purposes, we can use a bytes.Buffer, it has an implementation of oi.Writer
	var b bytes.Buffer
	car.Write(&b) // Here we call the write Method with our io.Writer
	fmt.Printf("%s",b.String() )
}
// Here we create our method that only takes a io.Writer as input
func (v *Vehicle) Write(w io.Writer) {
	b, _  := json.Marshal(*v)
	_, err := w.Write(b) // Here we write into the io.Writer
	if err != nil {
		return
	}
}
