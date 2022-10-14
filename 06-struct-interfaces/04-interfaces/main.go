package main

import "fmt"

type Animal interface {
	move() string
}

/* ***** Dog struct ***** */
type Dog struct{}

func (d *Dog) move() string {
	return "The dog walks"
}

/* ***** Fish struct ***** */
type Fish struct{}

func (f *Fish) move() string {
	return "The fish swims"
}

/* ***** Parrot struct ***** */
type Parrot struct{}

func (p *Parrot) move() string {
	return "the parrot sings"
}

/* ***** Help functions ***** */
func MoveAnimal(animal Animal) {
	fmt.Println(animal.move())
}

/* ***** Main function ***** */
func main() {
	dog := Dog{}
	MoveAnimal(&dog)

	fish := Fish{}
	MoveAnimal(&fish)

	parrot := Parrot{}
	MoveAnimal(&parrot)
}
