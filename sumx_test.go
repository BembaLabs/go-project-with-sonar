package main

import "testing"

func TestSumX(t *testing.T) {
	result := sumX(2, 3)
	expected := 7

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
