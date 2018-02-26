package main

import (
	"fmt"
	"math/rand"
	"time"
)

type honeyBee struct {
	name string
}

func (hb honeyBee) String() string {
	return hb.name
}

func (hb honeyBee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "buzzes about"
	default:
		return "flies to infinity and beyond"
	}
}

func (hb honeyBee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "pollen"
	default:
		return "nectar"
	}
}

type gopher struct {
	name string
}

func (g gopher) String() string {
	return g.name
}

func (g gopher) move() string {
	switch rand.Intn(2) {
	case 0:
		return "scurries along the ground"
	default:
		return "burrows in the sand"
	}
}

func (g gopher) eat() string {
	switch rand.Intn(5) {
	case 0:
		return "carrot"
	case 1:
		return "lettuce"
	case 2:
		return "radish"
	case 3:
		return "corn"
	default:
		return "root"
	}
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v eats the %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	rand.Seed(time.Now().UnixNano())

	animals := []animal{
		honeyBee{name: "Bzzz Lightyear"},
		gopher{name: "Go gopher"},
	}

	var sol, hour int

	for {
		fmt.Printf("%2d:00 ", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("The animals are sleeping.")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}

		time.Sleep(500 * time.Millisecond)

		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
}
