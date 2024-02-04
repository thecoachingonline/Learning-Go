package main

import (
	"fmt"
	"sort"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p:", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)

	var colors[3] string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[0])

	var number = [5] int{5, 3, 1, 2, 4}
	fmt.Println(number)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of number:", len(number))

	var colors1 = []string{"Red", "Green", "Blue"}
	fmt.Println(colors1)
	colors1 = append(colors1, "Purple")
	fmt.Println(colors1)

	colors1 = append(colors1[1:len(colors1)])
	fmt.Println(colors1)

	colors1 = append(colors1[:len(colors1)-1])
	fmt.Println(colors1)

	number1 := make([]int, 5)
	number1[0] = 134
	number1[1] = 72
	number1[2] = 32
	number1[3] = 12
	number1[4] = 156
	fmt.Println(number1)

	number1 = append(number1, 235)
	fmt.Println(number1)

	sort.Ints(number1)
	fmt.Println(number1)

	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	states["NY"] = "New York"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weigth)
	poodle.Weigth = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weigth)
}

type Dog struct {
	Breed string
	Weigth int
}