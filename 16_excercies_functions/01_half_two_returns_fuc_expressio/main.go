package main

import "fmt"

func main() {
	half := func(inNumber int) (int, bool) {
		var outBool bool

		if inNumber%2 == 0 {
			outBool = true
		} else {
			outBool = false
		}

		return inNumber / 2, outBool
	}

	fmt.Println(half(2))
}
