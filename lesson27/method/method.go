package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	p.age++
}

func main() {
	var nobody *person
	fmt.Println(nobody)

	nobody.birthday()
}
