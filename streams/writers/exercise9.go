package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("channels orchestrate mutexes serialize\n")
	proverbs.WriteString("cgo is not go\n")
	proverbs.WriteString("errors are values \n")
	proverbs.WriteString("don't panic\n")

	piper, pipew := io.Pipe()

	//write in writer end of pipe
	go func() {
		defer pipew.Close()
		io.Copy(pipew, proverbs)
	}()

	//read from reader end of pipe.
	io.Copy(os.Stdout, piper)
	piper.Close()

}
