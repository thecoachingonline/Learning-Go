package main

import (
	"fmt"
	"math/rand"
	"time"
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

	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(7) + 1
	// fmt.Println("Day", dow)

	var result1 string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result1 = "It's Sunday!"
		fallthrough
	case 2:
		result1 = "It's Monday!"
		fallthrough
	case 3:
		result1 = "It's some other day!"
	}
	fmt.Println(result1)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum < 200 {
			goto theEnd
		}
	}

	theEnd:
	fmt.Println("End of program")
}