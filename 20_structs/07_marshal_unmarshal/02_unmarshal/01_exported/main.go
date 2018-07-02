package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First, Last string
	Age         int
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Println("%T ", p1)

	bs := []byte(`{"First":"Juan","Last":"BG", "Age":21}`)

	json.Unmarshal(bs, &p1)

	fmt.Println("-----------------------------")

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Println("%T ", p1)

}
