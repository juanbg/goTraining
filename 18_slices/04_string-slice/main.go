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

	for i := 0; i < len(greeting); i++ {
		fmt.Println(greeting[i])
	}
}
