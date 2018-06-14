package main

import (
	"fmt"
)

func main() {

	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	n := avg(data...) //using variadic param or argunets we are sending to the func one value at the time.

	fmt.Println(n)

}

func avg(sf ...float64) float64 {
	var total float64

	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
