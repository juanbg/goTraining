package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for index := 0; index < 45; index++ {
		fmt.Println("Foo ", index)
	}
	wg.Done()

}

func bar() {
	for index := 0; index < 45; index++ {
		fmt.Println("Bar ", index)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
