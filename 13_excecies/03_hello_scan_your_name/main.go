package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("Could you write your name bub?")
	fmt.Scan(&name)

	fmt.Println("Hello World, and hello for u", name)
}
