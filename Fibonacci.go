package main

import "fmt"

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
