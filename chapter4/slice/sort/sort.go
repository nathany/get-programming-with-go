package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	sort.StringSlice(planets).Sort() // <1>
	fmt.Println(planets)             // <2>
}
