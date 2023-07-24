package main

import "fmt"

func zerovalur(ivalue int) {
	ivalue = 0
}
func zeropointer(ipointer *int){
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)

	zerovalur(i)
	fmt.Println("i from function zerovalur =", i)

	zeropointer(&i)
	fmt.Println("i value from function zeropointer =", i)
	fmt.Println("i address from function zeropointer =", &i)
}