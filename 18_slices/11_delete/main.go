package main

import (
	"fmt"
)

func main() {

	mySlice := []string{"Here", "is", "my", "fake", "love"}

	mySlice = append(mySlice[:3], mySlice[4:]...)

	fmt.Println(mySlice)
}
