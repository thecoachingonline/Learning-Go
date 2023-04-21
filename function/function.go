package main

import "fmt"

func hello(){
	fmt.Println("Hello")
}

func plus(value1 int, value2 int){
	result := value1 + value2
	fmt.Println("result = ", result)
}

func plus3value(value1, value2, value3 int) int{
	return value1 + value2 + value3
}

func main() {
	hello()
	plus(5, 4)
	result := plus3value(5, 5, 11)
	fmt.Println("result =", result)
}