package main

import (
	"fmt"
)

func main() {

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less then zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greatee then zero"
	}
	fmt.Println(result)

	if x := -42; x < 0 {
		result = "Less then zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater then zero"
	}
	fmt.Println(result)
}