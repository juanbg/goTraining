package main

import (
	"fmt"
)

const Hello = "hello"
const (
	A = iota
	B = iota
	C = iota
)

const (
	_  = iota
	BB = iota * 10
	CC = iota * 10
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func printKM() {
	fmt.Println("binary \t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
