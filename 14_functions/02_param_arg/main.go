package main

import (
	"fmt"
)

func main() {
	greet("John")
	greet("Josep")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

//func greet is declared with a parameter
//when calling greet, pass in an argument
