package main

import (
	"fmt"
)

func main() {
	fmt.Println(half(2))
}

func half(inNumber int) (int, bool) {

	var outBool bool

	if inNumber%2 == 0 {
		outBool = true
	} else {
		outBool = false
	}

	return inNumber / 2, outBool
}
