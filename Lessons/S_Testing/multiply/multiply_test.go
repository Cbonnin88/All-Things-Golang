package main

import (
	"fmt"
	"testing"
)

var ans int

func TestMultiply(t *testing.T) {
	answer := multiply(num1, num2)
	if answer != ans {
		t.Error("Expected:", ans, "Got:", answer)
	}
}

func TestTableMultiply(t *testing.T) {
	type test struct {
		input  []int
		answer int
	}
	tests := []test{
		{[]int{num1, num2}, ans},
	}
	for _, result := range tests {
		data := multiply(result.input...)
		if data != result.answer {
			t.Error("Expected:", result.answer, "Got:", data)
		}
	}
}

func ExampleMultiply() {
	fmt.Println(multiply(num1, num2))
	// Output:
	// 0
}

func BenchmarkMultiply(b *testing.B) {
	for m := 0; m < b.N; m++ {
		multiply(num1, num2)
	}
}
