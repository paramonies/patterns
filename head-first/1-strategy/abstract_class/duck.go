package abstract_class

import (
	"1-strategy/behavior"
	"fmt"
)

type AbstructDuck struct {
	FlyBehavior behavior.FlyBehavior
	QuackBehavior behavior.QuackBehavior

	Display func()
}

func (d *AbstructDuck) PerformFly() {
	d.FlyBehavior.Fly()
}

func (d *AbstructDuck) PerformQuack() {
	d.QuackBehavior.Quack()
}

func (d *AbstructDuck) Swim() {
	fmt.Println("All ducks float")
}

func (d *AbstructDuck) SetFlyBehavior(fb behavior.FlyBehavior) {
	d.FlyBehavior = fb
}

func (d *AbstructDuck) SetQuackBehavior(qb behavior.QuackBehavior) {
	d.QuackBehavior = qb
}