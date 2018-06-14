package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello", greet("juan", "bg"))
}

func greet(name, lName string) string {
	return fmt.Sprint(name, " ", lName)
}
