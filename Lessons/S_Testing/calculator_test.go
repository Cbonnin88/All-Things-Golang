package main

import "testing"

func TestCalculator(t *testing.T) {
	if Math(3) != 6 {
		t.Error("Expected 3 * 2 to equal 6")
	}
}
