package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Golang is a crappy language!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)

		fmt.Println("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
