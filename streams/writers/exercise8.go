package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./magic_msg.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	if _, err := io.WriteString(file, "Go is not that great;)"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
