package main

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	fmt.Println(pow)
	if len(pow) != 10  {
		t.Errorf("actual %q", pow)
	}
}
