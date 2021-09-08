package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	add := sum(100, 200)
	if add != 300 {
		t.Error("Expected:", 300, "Got:", add)
	}
}

func TestTableSum(t *testing.T) {
	type test struct {
		data    []int
		results int
	}
	tests := []test{
		{[]int{11, 11}, 22},
		{[]int{22, 22}, 44},
		{[]int{33, 44, 55}, 132},
		{[]int{1, 2, 3, 4}, 10},
	}

	for _, val := range tests {
		num := sum(val.data...)
		if num != val.results {
			t.Error("Expected:", val.results, "\n\t\tGot:", num)
		}
	}
}
