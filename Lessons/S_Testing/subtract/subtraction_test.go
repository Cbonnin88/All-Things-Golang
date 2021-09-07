package main

import "testing"

func TestMinus(t *testing.T) {
	subtract := minus(2000, 4444)
	if subtract != 2444 {
		t.Error("Expected:", 2444, "Got:", subtract)
	}
}
