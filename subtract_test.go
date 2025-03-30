package main

import "testing"

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	expected := 2

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
