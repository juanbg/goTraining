package main

import (
	"fmt"
)

func main() {
	greet("John", "BG")
	greet("Josep", "BG")
}

func greet(name string, lName string) {
	fmt.Println("Hello", name, lName)
}

//func greet is declared with a parameter
//when calling greet, pass in an argument
