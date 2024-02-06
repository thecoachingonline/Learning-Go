package main

import (
	"fmt"
)

func main()  {
	doSomething()
	sum := addValues(5, 8)
	fmt.Println("The sum is", sum)

	multiSum, multiCount := addAllValues(4, 7, 9, 45)
	fmt.Println("Sum of multiple values:", multiSum)
	fmt.Println("Count of items", multiCount)

	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()
	poodle.SpeakThreeTime()
	poodle.SpeakThreeTime()
}

func doSomething()  {
	fmt.Println("Doing something")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)

}

// Dog is a struct
type Dog struct {
	Breed string
	Weight int
	Sound string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTime() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}