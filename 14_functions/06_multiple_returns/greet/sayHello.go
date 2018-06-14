package greet

import (
	"fmt"
	"math/rand"
)

//CreateNickName func create a nickname compose by StringInt
func CreateNickName(name, lname string) (string, int) {
	nickName := fmt.Sprintf(
		"%s%d",
		combitanteStringsToGetNick(name, lname), rand.Int())

	return nickName, len(nickName)
}

func combitanteStringsToGetNick(name, lName string) (nickName string) {

	nickName = fmt.Sprint(
		string([]rune(name)[0:3]),
		string([]rune(lName)[0:2]))

	return nickName
}
