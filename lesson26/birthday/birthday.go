package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		age:        14,
		superpower: "imagination",
	}

	birthday(&rebecca)

	fmt.Printf("%+v\n", rebecca)
}
