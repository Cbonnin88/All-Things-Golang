package main

import (
	"fmt"
	"testing"
)

func TestExpression(t *testing.T) {
	saying := expression("Ethan Winters")
	if saying != "Rest while you can, Ethan Winters ,because I will hunt you, "+
		"and I will break you !!!" {
		t.Error("Got:", saying, "Expected", "Rest while you can, Ethan Winters ,"+
			"because I will hunt you,"+
			"and I will break you !!!")
	}
}

func ExampleExpression() {
	fmt.Println(expression("Ethan Winters"))
	// Output:
	// Rest while you can, Ethan Winters ,because I will hunt you, and I will break you !!!
}

func BenchmarkExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression("Ethan Winters")
	}
}
