package main

import (
	"fmt"
)

func facorial(x int) int {
	if x == 0 {
		fmt.Println("Time to die")
		return 1
	}
	fmt.Println(x)
	return x * facorial(x-1)
}

func main() {
	fmt.Println(facorial(4))
}
