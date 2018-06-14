package main

import (
	"fmt"
)

func main() {

	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	n := avg(data)

	fmt.Println(n)

}

func avg(sf []float64) float64 {
	var total float64

	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
