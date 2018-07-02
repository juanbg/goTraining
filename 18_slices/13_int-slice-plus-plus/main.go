package main

import "fmt"

func main() {
	/* 	mySlice := make([]int, 1)
	   	fmt.Println(mySlice[0])
	   	mySlice[0] = 7
	   	fmt.Println(mySlice[0])
	   	mySlice[0]++
	   	fmt.Println(mySlice[0]) */

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97,
	}

	var small int

	for i := 0; i < len(x); i++ {
		if i > 0 {
			if x[i-1] < x[i] {
				small = x[i-1]
			}
		}

	}

	fmt.Println(small)

}
