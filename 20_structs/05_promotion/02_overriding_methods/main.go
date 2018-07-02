package main

import (
	"fmt"
)

type Person struct {
	First  string
	Second string
	Age    int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm Just a Regular Person")
}

func (d DoubleZero) Greeting() {
	fmt.Println("Miss MoneyPenny, so god to see you")
}

func main() {
	p1 := Person{
		First:  "Juan",
		Second: "BG",
		Age:    21,
	}

	p2 := DoubleZero{
		Person: Person{
			First:  "Ara",
			Second: "Hg",
			Age:    21,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
}
