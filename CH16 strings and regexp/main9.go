package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for one person"

	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Splits >>" + x + "<<")
	}
}
