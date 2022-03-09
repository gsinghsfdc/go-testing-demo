package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 6
	actual := AddTwoNumbers(2, 4)

	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
