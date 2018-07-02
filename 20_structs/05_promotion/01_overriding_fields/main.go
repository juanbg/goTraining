package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type DobleZero struct {
	Person
	LicenseToKill bool
	Name          string
}

func main() {
	p1 := DobleZero{
		Person: Person{
			Name: "Tu padre",
			Age:  21,
		},
		LicenseToKill: true,
		Name:          "Juan",
	}
	fmt.Println(p1)
}
