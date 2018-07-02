package main

import (
	"fmt"
)

func main() {

	mapa := map[string]int{
		"Juan": 10,
		"Jose": 9,
		"Ara":  10 + 1,
	}
	for key := range mapa {
		fmt.Println(key)
	}

	for key, val := range mapa {
		fmt.Println(key, ": ", val)
	}

}

/**
Initialize a MAP using a composite literal where the key is a string and the value is an int; print out the map; range over the map printing out just the key; range over the map printing out both the key and the value

*/
