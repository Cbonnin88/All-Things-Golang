package main

import (
	"testing"
	"time"
)

func TestPackage(t *testing.T) {
	now := time.August
	if now != time.August {
		t.Error("Expected", time.August, "Got", now)
	}
}
