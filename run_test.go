package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	expected := []int{4, 2, 8, 6}
	results := Run()
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, but got %v", expected, results)
	}
}
