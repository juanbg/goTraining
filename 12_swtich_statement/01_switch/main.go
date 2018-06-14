package main

import (
	"fmt"
)

func main() {

	var name string

	fmt.Println("Please write your name to know who you are.")

	fmt.Scan(&name)

	switch name {
	case "Juan", "Beto":
		fmt.Println("Wassup Juan")
	case "Jose":
		fmt.Println("Wassup Jose")
	case "Ara":
		fmt.Println("Hello Darling")
		fallthrough
	case "Suegrita":
		fmt.Println("Hello you too!!")
		fallthrough
	case "Suegrito":
		fmt.Println("Of course for you too!")
	default:
		fmt.Println("Sorry bub, I don't know u.")
	}

	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}

}
