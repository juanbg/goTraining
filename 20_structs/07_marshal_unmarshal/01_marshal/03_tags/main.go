package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First  string
	Second string `json:"-"`
	Age    int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"JUAN", "BG", 21}

	bs, _ := json.Marshal(p1)

	fmt.Println(string(bs))
}
