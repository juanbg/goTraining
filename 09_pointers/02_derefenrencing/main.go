package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	/**
	* b is an int pointer
	* b points to the memory address where int is stored
	* to see  the value in that memory address, add * in front to the pointer in this case b
	* this is known as dereferencing
	* the * is an operator in this case
	 */
}
