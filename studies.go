package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("world 1")
	defer fmt.Println("world 2")
	defer fmt.Println("world 3")
	defer fmt.Println("world 4")
	defer fmt.Println("world 5")

	fmt.Println("hello")

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "nacl":
		fmt.Println("grap")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("piss %s.", os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Printf("Today %s.", today)
	case today + 1:
		fmt.Printf("Tomorrow %s.", today)
	case today + 2:
		fmt.Printf("In two days. %s.", today)
	case today + 3:
		fmt.Printf("In three days %s.", today)
	case today + 4:
		fmt.Printf("In four days. %s.", today)
	case today + 5:
		fmt.Printf("In five days %s.", today)
	case today + 6:
		fmt.Printf("In six days. %s.", today)
	default:
		fmt.Println("Too far away.")
	}

	pointers()

	//t := time.Now()
	//fmt.Printf("The current time is %s.", t)
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("Good Morning!")
	//case t.Hour() < 17:
	//	fmt.Println("Good Afternoon")
	//default:
	//	fmt.Println("Good Evening.")
	//
	//}

	fmt.Println("\n==========")
	fmt.Println(structure(12, 23).X)
	fmt.Println("==========")

	reachStructThroughPointer(234, 123)

	r, e := structLiterals(2, 1)
	fmt.Println(r, e)

}

func pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Println("pointers")
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(p) //prints memory address of the pointer
	fmt.Println("end pointers")

}

type Vertex struct {
	X int
	Y int
}

func structure(x, y int) Vertex {

	vert := Vertex{x, y}
	fmt.Println("\n==========")
	fmt.Println(vert)
	fmt.Println("==========")

	return vert
}

// You can access struct fields through a struct pointer.
func reachStructThroughPointer(x, y int) {
	v := structure(x, y)
	p := &v /// This is crap that I cannot get the address of the result of a function call

	p.X = 1e9
	fmt.Println(v)
}

func structLiterals(x, y int) (n int, err error) {
	v1 := Vertex{}
	v2 := Vertex{Y: 1}
	v3 := Vertex{}
	p := &Vertex{1, 2}
	return fmt.Println(v1, v2, v3, p)
}
