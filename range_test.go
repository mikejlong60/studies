package main

import (
	"fmt"
	"testing"
)

//When ranging over a slice, two values are returned for each iteration.
//First is the index, second is a copy of the element at that index.
//Its hard to write a test around this because of the imperative nature of loops.
//That's why you see no loops in FP.  Use recursion instead because you can test that easily.
// So I will just print the values.
func TestLookupTheValueInRangeByIndex(t *testing.T) {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
		fmt.Printf("%d\n", pow[i])
	}
}

func TestUseValueOnly(t *testing.T) {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

