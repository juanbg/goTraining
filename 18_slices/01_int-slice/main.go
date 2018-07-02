package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 45, 6, 7, 8, 9}
	/*
		fmt.Println(len(mySlice))
		fmt.Printf("%T %v\n", mySlice, &mySlice[1])
		fmt.Println(mySlice) */

	fmt.Println(mySlice)
	fmt.Println(mySlice[0:1])                   //slicing a slice
	fmt.Println(mySlice[2])                     //index access
	fmt.Println(string("This is a message"[6])) //index access a string is a slice of runes.
}

/*
a slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array

The value of an uninitialized slice is nil

Like arrays, slices are indexeable and have a length

the length of a slice can be discovered by the built-in function len
	-unlike arrays, silces are dynamics
		-the length may change during execution

the elements can be addressed by integer indices 0 throug len(s)-1

A slice, once initialized, is always associated with an underlying array taht holds its elements
	-its a reference type

The array underlying a slice may extend past the end of the slice

Make
	a new, initialized slice value for a given element type T, is made using the buit in function, make, which takes a sluce type and parameters specifyng the length and optionally the capacity
	a slice created with make, always allocates a new hidden array to which the returned slice values referes:
		make([]T, length, capacity)
		make[[]int, 50, 100]
*/
