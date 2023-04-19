package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =",product)

	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 25000
	fmt.Println("product =", product)

	delete(product, "iPad")
	fmt.Println("product =", product)

	product["iPhone"] = 20000
	fmt.Println("product =", product)

	value1 := product["iPhone"]
	fmt.Println("value1 =", value1)

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("courseName =", courseName)
}