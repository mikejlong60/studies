package main

import (
	"testing"
)

type Vertex2 struct {
	Lat, Long float64
}

func TestLookupTheValueInTheMap(t *testing.T) {
	var m map[string]Vertex2

	m = make(map[string]Vertex2)

	m["Bell Labs"] = Vertex2{40.123, -76.23141}

	actual := m["Bell Labs"]
	expected := Vertex2{40.123, -76.23141}

	if actual != expected {
		t.Errorf("actual %f , expected %f", actual, expected)
	}
}

func TestMapLiterals(t *testing.T) {

	var m = map[string]Vertex2{"google": Vertex2{40.123, -76.23141}, "fred": Vertex2{22.123, -76.23141}}

	actual := m["fred"]
	expected := Vertex2{22.123, -76.23141}

	if actual != expected {
		t.Errorf("actual %f , expected %f", actual, expected)
	}
}

func TestMapLiteralsOmitTypeName(t *testing.T) {

	var m = map[string]Vertex2{"google": {40.123, -76.23141}, "fred": {22.123, -76.23141}}

	actual := m["fred"]
	expected := Vertex2{22.123, -76.23141}

	if actual != expected {
		t.Errorf("actual %f , expected %f", actual, expected)
	}
}

func TestMapMutations(t *testing.T) {
	var m = make(map[string]int)

	m["fred"] = 42
	actual := m["fred"]
	expected := 42
	if actual != expected {
		t.Errorf("actual %d , expected %d", actual, expected)
	}

	//Returns default monoidal value if its not there
	actual = m["fred1"]
	expected = 0
	if actual != expected {
		t.Errorf("actual %d , expected %d", actual, expected)
	}

	//If you grab the error code it will tell you if it is not there. Test that a key is present with a two-value assignment.
	actual, ok := m["fred1"]
	expectedok := false
	if ok != expectedok {
		t.Errorf("actual %t , expected %t", ok, expectedok)
	}

	//Does nothing if element does not exist.
	delete(m, "fred1")
	expected = 0
	if actual != expected {
		t.Errorf("actual %d , expected %d", actual, expected)
	}

}
