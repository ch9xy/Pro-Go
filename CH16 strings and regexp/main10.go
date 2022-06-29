package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "This  is  double  space"

	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Splits >>" + x + "<<")
	}
}
