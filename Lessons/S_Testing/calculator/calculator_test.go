package main

import "testing"

func TestCalculator(t *testing.T) {
	if math(3) != 6 {
		t.Error("Expected 3 * 2 to equal 6")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{3, 6},
		{6, 12},
		{9, 18},
		{2, 4},
		{0, 0},
		{1, 2},
		{10, 20},
	}

	for _, test := range tests {
		if output := math(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, received:{}", test.input, test.expected, output)
		}
	}
}
