package main

import (
	"fmt"
)

const metersToYars float64 = 1.09361

func main() {
	var meters float64
	fmt.Println("How many meters did you swim?")
	fmt.Scan(&meters)
	yards := meters * metersToYars
	fmt.Println(meters, " meters is ", yards, " yards.")

}
