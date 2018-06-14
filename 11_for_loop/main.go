package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("this is my message number: ", i)
	}

	ii := 0
	for ii < 10 {
		fmt.Print("ii is not bigger than 10")
		ii++
	}

	/*	for {
			fmt.Println("Infinite")
		}
	*/
	a := 0
	for {
		fmt.Println(a)

		if a >= 10 {
			break
		}
		a++
	}

	j := 0

	for {
		j++
		if j%2 == 0 {
			continue // if condition is true it ends this loop and start with new one
		}

		fmt.Println(j)

		if j >= 50 {
			break
		}
	}
}
