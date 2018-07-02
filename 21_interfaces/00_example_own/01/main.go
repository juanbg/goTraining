package main

import (
	"fmt"
)

/*
Quester is about any quest
*/
type Quester interface {
	QuestStatement()
}

/*
KillDragonQuest implementation of Quester
*/
type KillDragonQuest struct {
	Statement string
	Duration  float32
	Complete  bool
}

/*
QuestStatement is a method od KillDragonQuest
*/
func (q KillDragonQuest) QuestStatement() {
	fmt.Println(q.Statement)
}

/*
Knighter Generic Knight
*/
type Knighter interface {
	Embark()
}

/*
Embark implementation
*/
func (k Knight) Embark() {
	k.Quester.QuestStatement()
}

/*
Knight is Specialiced Knight
*/
type Knight struct {
	Name   string
	Age    string
	Enable bool
	Quester
}

func main() {

	/* 	var ques = new(KillDragonQuest)
	 */
	ques := KillDragonQuest{"asd", 2.2, true}
	Knight := Knight{"Juan", "12", true, ques}

	Knight.Embark()

}
