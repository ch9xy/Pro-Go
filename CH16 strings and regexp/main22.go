package main

import (
	"fmt"
	//"strings"
	"regexp"
)

func main() {
	pattern, compilerErr := regexp.Compile("[A-z]oat")

	description := "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"

	if compilerErr == nil {
		fmt.Println("Description: ", pattern.MatchString(description))
		fmt.Println("Question: ", pattern.MatchString(question))
		fmt.Println("Preference: ", pattern.MatchString(preference))
	}
}
