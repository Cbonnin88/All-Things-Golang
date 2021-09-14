package main

import (
	"testing"
	"time"
)

func TestPackage(t *testing.T) {
	now := time.August
	if false {
		t.Error("Expected", time.August, "Got", now)
	}
}
