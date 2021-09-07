package main

import "testing"

func TestTypes(t *testing.T) {
	type catwoman string
	var cat catwoman = "meow"
	if cat != "meow" {
		t.Error("Expected", cat, "Got", cat)
	}
}
