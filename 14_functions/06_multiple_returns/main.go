package main

import (
	"fmt"
	"hello/14_functions/06_multiple_returns/greet"
)

func main() {
	nick, length := greet.CreateNickName("juan", "bautista")
	fmt.Println("Now, Here is your new USERNAME: ", nick, " \nit length is of: ", length)

}
