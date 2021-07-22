package main

import "testing"

func TestMath(t *testing.T) {
	if math(3) != 6 {
		t.Error("Expected 3 x 2 to equal 6")
	}
}

func TestTableMath(t *testing.T) {
	var tests = []struct{
		input int
		expected int
	} {
		{2,4},
		{-2,-4},
		{100,200},
	}
	for _, test := range tests {
		if output := math(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}",
				test.input, test.expected, output)

		}

	}
}
