package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)  //43
	fmt.Println(&a) //0xc4200140d8

	var b *int = &a
	fmt.Println(b)  // 0xc4200140d8
	fmt.Println(*b) //43

	*b = 42 //b says, "The value  at this address, cange to 42"

	fmt.Println(a) //42

	/*
		this is useful
		we can pass memory address instead of a bunch of values (we can pass a reference)
		and then we can still change the value whatever is stored at that memory address
		this makes our program more efficient
		we dont have to pass around large amounts of data
		we only have to pass around address

		Everything is PASS BY VALUE in go, btw
		when we pass memory address, we are passing a value
	*/
}
