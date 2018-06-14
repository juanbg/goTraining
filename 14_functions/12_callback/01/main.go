package main

import (
	"fmt"
)

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		/* fmt.Println(n) */
	})

	visit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(n int) {
		n = n * n

		fmt.Println(n)
	})
}

// Callback is passing a func as an argument
