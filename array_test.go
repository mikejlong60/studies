package main

import "testing"

func TestArrayEqualityBasedUponValue(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	var e [2]string
	e[0] = "Hello"
	e[1] = "World"
	if a != e {
		t.Errorf("actual %q , expected %q", a, e)
	}
}


func TestArraySlicing(t *testing.T) {
	a  := [5]int{1,2,3,4,5}

	var a1 = a[1:4]

	if len(a1) != 3 {
		t.Error("expected slice to be a 3 element array")
	}
}
