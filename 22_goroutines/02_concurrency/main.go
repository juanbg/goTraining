package main

import (
	"fmt"
)

func foo() {
	for index := 0; index < 45; index++ {
		fmt.Println("foo ", index)
	}
}

func bar() {
	for index := 0; index < 45; index++ {
		fmt.Println("bar ", index)
	}
}

func main() {
	go foo()
	go bar()
}
