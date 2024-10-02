package main

import (
	"fmt"
)

func main() {
	var quantity int
	quantity = 4

	var length, width float64 = 1.2, 2.4
	var customerName = "Damon"

	quantity = 5

	var zero int
	var falseBool bool
	var emptyString string

	fmt.Println(quantity)
	fmt.Println(length * width)
	fmt.Println(customerName)
	fmt.Println(zero)
	fmt.Println(falseBool)
	fmt.Println(emptyString)
}
