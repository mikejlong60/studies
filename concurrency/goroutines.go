package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5000; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Sprintf("%v ---  %v", s, i)
	}
}

func main() {
	for i := 0; i < 50000000; i++ {

		go say(fmt.Sprintf("world%v", i))
		fmt.Println(i)
	}
	say("hello")
}
