package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("This is a bad practice")
	}()
}
