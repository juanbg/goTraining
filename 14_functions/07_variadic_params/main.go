package main

import (
	"fmt"
)

func main() {
	n := avg(1, 44, 2, 213, 2, 1, 45, 4, 12, 3, 12)
	fmt.Println(n)
}

func avg(sf ...float64) float64 {
	var total float64
	for _, v := range sf { // it is used to iterate a slice or map the first return is an index, the second one is the value
		total += v
	}

	return total / float64(len(sf))
}
