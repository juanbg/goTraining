package main

import (
	"fmt"
)

func main() {
	customerNumber := make([]int, 3)
	//3 is the length and capacity
	//length - number of elements referred to by the slice
	//capacity - number of elements in the underlying array

	customerNumber[0] = 7
	customerNumber[1] = 5
	customerNumber[2] = 8

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	//3 is length - number of elements referred to by the slice
	//5 is capacity - number of elements in the underlying array

	greeting[0] = "Hola"
	greeting[1] = "ホア"
	greeting[2] = "привет"

	greeting = append(greeting, "Ciao")

	fmt.Println(greeting)

}
