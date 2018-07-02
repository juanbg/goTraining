package main

import (
	"fmt"
)

func main() {
	greeting := []string{
		"Hello",
		"Hola",
		"ホア",
		"ciao",
		"привет",
	}

	fmt.Println(greeting[:])
	fmt.Println(greeting[:3])
	fmt.Println(greeting[3:])
	fmt.Println(greeting[1:4])
}
