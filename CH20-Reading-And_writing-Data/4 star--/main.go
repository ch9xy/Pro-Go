package main

import (
	"io"
	//"strings"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)
}
