package main

import (
	"testing"
)

func TestStructDefaults(t *testing.T) {
	cases := []struct {
		actual, expected Vertex
	}{
		{Vertex{1, 2}, Vertex{1, 2}},
		{Vertex{Y: 1}, Vertex{0, 1}},
		{Vertex{}, Vertex{0, 0}},
	}
	for _, c := range cases {
		if c.actual != c.expected {
			t.Errorf("%q, want %q", c.actual, c.expected)
		}
	}
}

func TestStructAccess(t *testing.T) {
	v := Vertex{1, 2}
	if v.Y != 2 {
		t.Errorf("actual %q , expected %q", v.Y, 2)

	}
}


func TestStructAccessWithPointerSyntax(t *testing.T) {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e3
	expected := Vertex{1000, 2}
	if v != expected {
		t.Errorf("actual %q , expected %q", v, expected)
	}
}
