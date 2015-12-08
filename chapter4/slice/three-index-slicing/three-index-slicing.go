package main

import "fmt"

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	// tag::main[]
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4] // <1>
	worlds := append(terrestrial, "Ceres")
	// end::main[]
	fmt.Println(planets)

	dump("terrestrial", terrestrial)

	// tag::two[]
	terrestrial = planets[0:4] // <1>
	worlds = append(terrestrial, "Ceres")

	fmt.Println(planets) // <2>
	// end::two[]

	dump("terrestrial", terrestrial)

	_ = worlds
}
