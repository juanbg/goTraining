package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 43, 12, 3, 345, 1}

	for _, s := range slice {
		fmt.Println(s)
	}

	for index, v := range slice {
		fmt.Println(index, ".- ", v)
	}
}

/*
Initialize a SLICE of int using a composite literal; print out the slice; range over the slice printing out just the index; range over the slice printing out both the index and the value
*/
