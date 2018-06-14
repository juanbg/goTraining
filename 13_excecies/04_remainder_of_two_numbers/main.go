package main

import (
	"fmt"
)

func main() {

	var smallNumber, largerNumber int

	fmt.Println("Can you write a number?")
	fmt.Scan(&smallNumber)
	fmt.Println("Good \n, now, Can you write a new number, but this time largest?")
	fmt.Scan(&largerNumber)

	if largerNumber > smallNumber {
		fmt.Printf("here is the remainder %d", largerNumber%smallNumber)
	} else {
		fmt.Println("Ha ha ha, your sencod number is smallest")
	}

}
