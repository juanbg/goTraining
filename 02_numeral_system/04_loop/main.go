package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}

	/*	sum := 1
		for sum < 1000 {
			sum += sum
			fmt.Println(sum)
		}*/
}
