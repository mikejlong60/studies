package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/user/studies/streams/readers"
)

func main() {
	reader := readers.NewAlphaReader(strings.NewReader("hello! it's 9am, where is the sun?"))
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)

		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
	fmt.Println()

}
