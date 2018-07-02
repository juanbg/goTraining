/**
The idea for this Kata came from 9gag today.here

You'll have to translate a string to Pilot's alphabet (NATO phonetic alphabet) wiki.

Like this:

Input: If you can read

Output: India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta

Some notes

Keep the punctuation, and remove the spaces.
Use Xray without dash or space.
**/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%s%s%s", "-", ToNato(""), "-")
}

/*
ToNato tranlate a word to NATO
*/
func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

	var natoTranslated string

	eachCharacter := []rune(words)

	for _, v := range eachCharacter {
		if v >= 33 && v <= 65 {
			natoTranslated += fmt.Sprint(string(v), " ")

		}
		for i := 0; i < len(nato); i++ {
			if strings.EqualFold(string(v), string([]rune(nato[i])[0:1])) {
				natoTranslated += fmt.Sprint(nato[i], " ")
			}
		}

	}

	return strings.Replace(string([]rune(natoTranslated)[0:len(natoTranslated)-1]), "-", "", -1)

}
