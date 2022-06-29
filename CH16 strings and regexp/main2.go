package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	price := "E100"
	fmt.Println("Strings Prefix:", strings.HasPrefix(price, "E"))
	fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price),
		[]byte{226, 130}))
}
