package main

import (
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
		test{[]int{num1, num2}, ans},
		test{[]int{num1, num2}, ans},
		test{[]int{num1, num2}, ans},
		test{[]int{num1, num2}, ans},
		test{[]int{num1, num2}, ans},
	}
	for _, result := range tests {
		data := multiply(result.input...)
		if data != result.answer {
			t.Error("Expected:", result.answer, "Got:", data)
		}
	}

}
