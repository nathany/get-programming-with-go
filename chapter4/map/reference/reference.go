package main

import "fmt"

func main() {
	// tag::main[]
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planetsMarkII := planets
	planets["Earth"] = "💥"

	fmt.Println(planets)       //<1>
	fmt.Println(planetsMarkII) //<1>

	delete(planets, "Earth")   //<2>
	fmt.Println(planetsMarkII) //<3>
	// end::main[]
}
