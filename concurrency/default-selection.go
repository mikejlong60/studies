package main

import "fmt"
import "time"

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(10000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("       .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
