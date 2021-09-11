package main

import (
	"1-strategy/abstract_class"
	"1-strategy/behavior"
	"fmt"
)

type MallardDuck struct {
	*abstract_class.AbstructDuck
}

func NewMallrdDuck() *MallardDuck {
	d := &abstract_class.AbstructDuck{
		FlyBehavior: &behavior.FlyWithWings{},
		QuackBehavior: &behavior.Quack{},
		Display: func() {
			fmt.Println("I'm a real MallardDuck!")
		},
	}
	return &MallardDuck{d}
}

type ModelDuck struct {
	*abstract_class.AbstructDuck
}

func NewModelDuck() *ModelDuck {
	d := &abstract_class.AbstructDuck{
		FlyBehavior: &behavior.FlyNoWay{},
		QuackBehavior: &behavior.Quack{},
		Display: func() {
			fmt.Println("I'm a modil duck!")
		},
	}
	return &ModelDuck{d}
}

func main() {
	mallardDuck := NewMallrdDuck()
	mallardDuck.Display()
	mallardDuck.PerformFly()
	mallardDuck.PerformQuack()
	mallardDuck.Swim()

	fmt.Println("======================")

	modelDuck := NewModelDuck()
	modelDuck.Display()
	modelDuck.PerformFly()
	modelDuck.PerformQuack()
	modelDuck.Swim()

	fmt.Println("======================")

	modelDuck.SetFlyBehavior(&behavior.FlyRocketPowered{})
	fmt.Println("New Flying Behaviour of Model Duck:")
	modelDuck.PerformFly()
}