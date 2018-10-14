package main

import "testing"

func TestAdder(t *testing.T) {
	result := Adder(4, 7)
	if result != 11 {
		t.Fatal("4+7 did not equal 11")
	}
}

