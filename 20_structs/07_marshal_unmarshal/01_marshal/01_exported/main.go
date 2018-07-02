package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Second      string
	Age         int
	notExported int
}

func (p Person) Greeting() {
	fmt.Println("Hello ")
}

func main() {

	p1 := Person{"Juan", "BG", 21, 00}
	p1.Greeting()
	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T %v", bs, bs)

	fmt.Println(string(bs))
}
