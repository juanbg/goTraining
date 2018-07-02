package main

import (
	"fmt"
)

/*
Knighter Generic Knight
*/
type Knighter interface {
	Embark()
}

/*
Quester is about any quest
*/
type Quester interface {
	QuestStatement()
}

/*
KillDragonKnight is Specialiced Knight
*/
type KillDragonKnight struct {
	Name   string
	Age    string
	Enable bool
}

/*
SaveDamsel specialiced Knight
*/
type SaveDamsel struct {
	Name   string
	Age    string
	Enable bool
}

/*
Embark implementation for SaveDamesel
*/
func (k SaveDamsel) Embark() {
	fmt.Println("I'll give my life for you madame!")
}

/*
Embark implementation for KillDragonKnight
*/
func (k KillDragonKnight) Embark() {
	fmt.Println("I'm gona kill you damn dragon!")
}

func main() {

	var k Knighter

	k = new(KillDragonKnight)

	k.Embark()

	k = new(SaveDamsel)

	k.Embark()
}
