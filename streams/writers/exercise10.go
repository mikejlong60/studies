package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\b') //('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("found EOF before finding delimiter")
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Println(line)
	}

}
