package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y--
}

func (t *turtle) down() {
	t.y++
}

func (t *turtle) left() {
	t.x--
}

func (t *turtle) right() {
	t.x++
}

func main() {
	var t turtle
	t.up()
	t.up()
	t.left()
	t.left()
	fmt.Println(t)
	t.down()
	t.down()
	t.right()
	t.right()
	fmt.Println(t)
}
