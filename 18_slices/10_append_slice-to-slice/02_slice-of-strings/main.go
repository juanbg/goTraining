package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"Monday", "Tuesday", "Wednesday"}
	myOtherSlice := []string{"Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}
