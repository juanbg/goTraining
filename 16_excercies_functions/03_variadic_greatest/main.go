package main

import (
	"fmt"
)

func main() {
	fmt.Println(theGreatest(10, 1, 2, 1, 3))
}

func theGreatest(inComingNumbers ...int) int {
	var bigger int

	for _, v := range inComingNumbers {
		for a := 0; a < len(inComingNumbers); a++ {
			if v < inComingNumbers[a] {
				bigger = inComingNumbers[a]
			}
		}
	}

	return bigger
}
