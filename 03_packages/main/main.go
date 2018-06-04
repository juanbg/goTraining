package main

import (
	"fmt"
	"hello/03_packages/utils"
)

func main() {
	fmt.Printf("%s, how are you?", utils.Myname)
	fmt.Printf(utils.Reverse("\nThis is a text\n"))
}
