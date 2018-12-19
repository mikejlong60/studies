package main

import "fmt"

//This is bad because it uses mutation which means lots of good things can't happen.  Can't distribute, can't do in parallel, can't compose, etc.  We should not be writing programs in this style.
func fibonacci() func() int {

	var fst = 0
	var snd = 1  
	return func() int {

		next := fst + snd
		fst = snd
		snd = next
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
