package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Juan", "Bautista Grande"))
}

func greet(name, lName string) (s string) {
	s = fmt.Sprint(name, " ", lName)

	return
}
