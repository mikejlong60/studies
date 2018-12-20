package main

import (
	"fmt"
	"testing"
)

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
	a := [5]int{1, 2, 3, 4, 5}

	var a1 = a[1:4]

	if len(a1) != 3 {
		t.Error("expected slice to be a 3 element array")
	}
}

func TestSlicingCreatesReferencesThatYouCanActuallyMutateHowMoronic(t *testing.T) {
	original := [4]string{"john", "paul", "ringo", "george"}

	a := original[0:2]
	a[0] = "XXX"

	if original[0] != "XXX" {
		t.Errorf("actual %q , expected %q", a[0], "XXX")
	}
}

func TestSlicesLiteralsAreArrayLiteralsWithoutLength(t *testing.T) {
	q := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(q)

	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{11, true},
	}
	if len(s) != 4 {
		t.Errorf("actual %q, expected %q", len(s), 4)
	}

}

func TestSliceDefaults(t *testing.T) {
	q := []int{2, 3, 4, 5, 6, 7}

	c := q[0:6]
	d := q[:6]
	e := q[0:]
	f := q[:]
	if (len(c) != len(q)) || (len(d) != len(q)) || (len(e) != len(q)) || (len(f) != len(q)) {
		t.Errorf("actual %q %q %q %q, expected %q", c, d, e, f, q)
	}
}

func TestSliceLength(t *testing.T) {
	q := []int{2, 3, 4, 5, 6, 7}
	a := q[:0]
	b := q[:4]
	c := q[2:]

	if len(a) != 0 || len(b) != 4 || len(c) != 4 {
		t.Errorf("actual %q %q %q, expected: 0, 4 ,4", len(a), len(b), len(c))
	}
}

func TestNilSlice(t *testing.T) {
	var q []int

	if q != nil || cap(q) != 0 {
		t.Errorf("actual %q, expected: 0, 4 ,4", q)
	}
}

func TestMakingSlice(t *testing.T) {
	q := make([]int, 5)

	if len(q) != 5 || cap(q) != 5 {
		t.Errorf("actual %q, expected: 0, 4 ,4", q)
	}
}

func TestSlicesOfSlices(t *testing.T) {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	if len(board) != 3 || board[0][2] != "X" {
		t.Errorf("actual %q, expected: 0, 4 ,4", board)
	}
}

func TestAppendingToSlice(t *testing.T) {
	var s []int
	s = append(s, 0)
	s = append(s, 1)
	s = append(s, 3, 4, 5)

	if len(s) != 5 || s[1] != 1 {
		t.Errorf("actual %q", s)
	}
}
