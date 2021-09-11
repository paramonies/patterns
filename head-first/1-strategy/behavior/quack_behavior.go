package behavior

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct {}
func (q *Quack) Quack() {
	fmt.Println("Quack")
}

type Squeak struct {}
func (q *Squeak) Quack() {
	fmt.Println("Squeak")
}

type MuteQuack struct {}
func (q *MuteQuack) Quack() {
	fmt.Println("<< Silent >>")
}