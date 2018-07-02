package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{"Juan", "BG", 21}
	p2 := person{"José", "BG", 18}

	p1.firstName = "Juan Alberto"

	fmt.Println("Hello ", p1.firstName, ",", p2.firstName)
}
