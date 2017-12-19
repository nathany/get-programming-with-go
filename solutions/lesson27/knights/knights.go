package main

import (
	"fmt"
)

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up a %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) giveTo(recipient *character) {
	if c == nil || recipient == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
		return
	}
	if recipient.leftHand != nil {
		fmt.Printf("%v's hands are full\n", recipient.name)
		return
	}
	recipient.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v's gives %v a %v\n", c.name, recipient.name, recipient.leftHand.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v is carrying nothing", c.name)
	}
	return fmt.Sprintf("%v is carrying a %v", c.name, c.leftHand.name)
}

func main() {
	arthur := &character{name: "Arthur"}

	shrubbery := &item{name: "shrubbery"}
	arthur.pickup(shrubbery)

	knight := &character{name: "Knight"}
	arthur.giveTo(knight)

	fmt.Println(arthur)
	fmt.Println(knight)
}
