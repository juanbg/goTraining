package main

import (
	"fmt"
)

func main() {
	greet := makeAGreeter()
	fmt.Println(greet())
}

func makeAGreeter() func() string {
	return func() string {
		return "Hello World"
	}

}
