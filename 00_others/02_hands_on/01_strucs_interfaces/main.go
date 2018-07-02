package main

import (
	"fmt"
	"math"
)

/************************************/

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

/*************************************/

type circle struct {
	radious float64
}

func (c circle) area() float64 {
	return (c.radious * math.Pi) * (c.radious * math.Pi)
}

/***********************************/

type shape interface {
	area() float64
}

func main() {
	c := circle{64}
	s := square{2}

	fmt.Println(info(c))
	fmt.Println(info(s))
}

func info(s shape) interface{} {

	return s.area()
}
