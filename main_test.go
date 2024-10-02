package main

import (
	"testing"
)

func TestNameIsEvolution(t *testing.T) {
	expected := "Evolution"
	actual := name()
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
