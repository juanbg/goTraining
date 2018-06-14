package main

import (
	"fmt"
)

func filter(numbers []int, callback func(int) bool) []int { //we construct a function that needs a slice of ints, and a function which accept like argument an int and returns a bool and the function returns an int
	xs := []int{} //different from xs in main func

	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	} /*  */

	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)
}
