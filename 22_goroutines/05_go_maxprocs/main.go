package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo() {
	for index := 0; index < 45; index++ {
		fmt.Println("Foo ", index)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()

}

func bar() {
	for index := 0; index < 45; index++ {
		fmt.Println("Bar ", index)
		time.Sleep(time.Duration(3 * time.Millisecond))

	}
	wg.Done()
}

func init() {
	fmt.Println("You will use ", runtime.NumCPU(), " processor")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
