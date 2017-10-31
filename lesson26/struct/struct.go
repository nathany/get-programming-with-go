package main

import "fmt"

func main() {
	type person struct {
		name, superpower string
		age              int
	}

	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	timmy.superpower = "flying"

	fmt.Printf("%+v\n", timmy)
}
