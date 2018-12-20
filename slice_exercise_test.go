package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"testing"
)

//When ranging over a slice, two values are returned for each iteration.
//First is the index, second is a copy of the element at that index.
//Its hard to write a test around this because of the imperative nature of loops.
//That's why you see no loops in FP.  Use recursion instead because you can test that easily.
// So I will just print the values.
func TestSliceExercise(t *testing.T) {
	xaxis := make([]uint8, 10)
	yaxis := make([]uint8, 10)
	for i := range xaxis {
		xaxis[i] = 1 << uint(i)
		fmt.Printf("%d\n", xaxis[i])
	}
	for i := range yaxis {
		yaxis[i] = 1 << uint(i)
		fmt.Printf("%d\n", yaxis[i])
	}

}

func Pic(dx, dy int) [][]uint8 {
	fmt.Printf("dude 1: %d %d\n", dx, dy)
	xaxis := make([]uint8, dx, dx)
	for i := 0; i < dx; i++ {
		xval := uint8(dx ^ dy) //uint8(i * dy)//dx ^ dy)//(dx+dy)/2)//uint8(i)
		fmt.Printf("dude 2: %d\n", i)
		fmt.Printf("dude 21: %d\n", xval)
		xaxis[i] = xval // uint8(i * dy)//dx ^ dy)//(dx+dy)/2)//uint8(i)
	}

	result := make([][]uint8, dy, dy)
	for i := 0; i < dy; i++ {
		result[i] = xaxis
	}

	return result
}

func TestPic(t *testing.T) {
	pic.Show(Pic)
}
