package main

import {
	"fmt"
}

const aConst int = 64

func main() {

	var aString string = "This is Go!"

	fmt.Println(aString)
	fmt.Println("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	ver anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Println("The variable's type is %T\n", anotherString)

	ver anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Println("The variable's type is %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Println("The variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Println("The variable's type is %T\n", aConst)
}