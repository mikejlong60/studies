package main

import (
	"fmt"
	"io"
	"os"

	"github.com/user/studies/streams/readers"
)

func main() {
	file, err := os.Open("./exercise3.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := readers.NewAlphaReader(file)
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
