package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 6
	actual := AddTwoNumbers(2, 4)

	if actual != expected {
		t.Fatal(fmt.Sprintf("expected %v, got %v", expected, actual))
	}
}
