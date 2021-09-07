package main

import "testing"

func TestHello(t *testing.T) {
	hello := "Go Testing"
	if hello != "Go Testing" {
		t.Error("Expected", hello, "Got", hello)
	}
}
