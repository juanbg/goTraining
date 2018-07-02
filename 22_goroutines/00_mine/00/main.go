package main

import (
	"fmt"
	"time"
)



func main()  {
	go loadingMessage()
	fmt.Println("\r",fnumber(8))

}

func fnumber(tNumber int) int{

	if tNumber < 2 {
		return tNumber
	}

	return fnumber(tNumber-1) + fnumber(tNumber-2)
}


func loadingMessage(){
	for {
		for _, value := range []string  {"Printing","Printing.","Printing..","Printing..."}{
			time.Sleep(10 * time.Millisecond)
			fmt.Print("\r",value)

		}
	}
}