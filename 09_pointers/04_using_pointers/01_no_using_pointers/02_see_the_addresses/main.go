package main

import (
	"fmt"
)

func zero(z int) {
	fmt.Printf("%p\n", &z) //address in fuc zero
	fmt.Println(&z)
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) //Adress on func main
	fmt.Println(&x)        //address in main

	zero(x)
	fmt.Println(x) //still 5
}
