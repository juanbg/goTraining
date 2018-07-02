package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first  string
	second string
	age    int
}

func main() {
	p1 := Person{"Juan", "BG", 21}
	fmt.Println(p1)

	bs, _ := json.Marshal(p1)

	fmt.Println(string(bs))
}
