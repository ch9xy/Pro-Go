package main

import (
	"regexp"
	//"fmt"
)

func main() {
	pattern := regexp.MustCompile(
		"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")

}
