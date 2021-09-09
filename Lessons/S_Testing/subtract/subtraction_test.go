package main

import (
	"fmt"
	"testing"
)

func TestMinus(t *testing.T) {
	subtract := minus(2000, 4444)
	if subtract != 2444 {
		t.Error("Expected:", 2444, "Got:", subtract)
	}
}

func TestTableMinus(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		{[]int{2, 5}, 3},
		{[]int{2, 4}, 2},
		{[]int{100, 9999}, 9899},
		{[]int{12, 65}, 53},
	}
	for _, result := range tests {
		input := minus(result.data...)
		if input != result.answer {
			t.Error("Expected:", result.answer, "Got:", input)
		}
	}
}

func ExampleSubtraction() {
	fmt.Println(minus(12, 100))
	// Output:
	// 88
}

func BenchmarkSubtraction(b *testing.B) {
	for s := 0; s < b.N; s++ {
		minus(12, 100)
	}
}
