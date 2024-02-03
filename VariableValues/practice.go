package main

import (
	"fmt"
	"bufio"
	"math"
	"os"
	"strconv"
	"time"
)

const aConst int = 64

func main() {

	var aString string = "This is Go!"

	fmt.Println(aString)
	fmt.Println("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Println("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Println("The variable's type is %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Println("The variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Println("The variable's type is %T\n", aConst)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(string.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", aFloat)
	}

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Println("Circumference: %.2f\n", circumference)

	n := time.Now()
	fmt.Println("I recorded this video at ", n)

	t := time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Println("The type of parsedTime is %T\n", parsedTime)

}