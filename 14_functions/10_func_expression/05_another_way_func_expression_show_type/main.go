package main

import (
	"fmt"
)

func main() {
	greet := makeAGreeter()
	fmt.Println(greet())
	fmt.Printf("%T", greet)
}

func makeAGreeter() func() string {
	return func() string {
		return "Hello World"
	}

}
