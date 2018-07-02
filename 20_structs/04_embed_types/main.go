package main

import (
	"fmt"
	"hello/20_structs/04_embed_types/inside"
)

func main() {

	p1 := inside.DobleZero{
		Person: inside.Person{
			FirstName:  "Juan",
			SecondName: " BG",
			Age:        21,
		},
		LicenseToKill: true,
	}

	fmt.Println(p1)
}
