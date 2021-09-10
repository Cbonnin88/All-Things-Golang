package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	helloWorld := "TestA"
	if helloWorld != "TestA" {
		t.Error("Expected:", "TestA", "Got:", helloWorld)
	}
}
