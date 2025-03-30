package main

import "testing"

func TestMutiply(t *testing.T) {
	result := multiply(4, 2)
	expected := 8

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
