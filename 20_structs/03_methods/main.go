package main

import (
	"fmt"
)

type person struct {
	first  string
	second string
	age    int
}

func (p person) giveFullyName() string {
	return fmt.Sprint(p.first, p.second)
}

func main() {
	p1 := person{"Juan ", "Bg", 21}
	p2 := person{"Ara ", "Hg", 20}

	fmt.Println(p1.giveFullyName())
	fmt.Println(p2.giveFullyName())
}
