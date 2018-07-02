package main

import (
	"fmt"
	"time"
)

type foo int

func main() {
	var myAge foo
	myAge = 21
	fmt.Printf("%T %v\n", myAge, myAge)

	yourAge := 44
	fmt.Printf("%T %v\n", yourAge, yourAge)

	fmt.Println(myAge + foo(yourAge)) // you need convert

	t0 := time.Now()

	fmt.Println(t0)
}
