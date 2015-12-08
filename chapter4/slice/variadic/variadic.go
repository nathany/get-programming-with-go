package main

import "fmt"

// tag::terraform[]
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds)) //<1>

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

// end::terraform[]

func main() {
	// tag::main[]
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds) // <1>
	// end::main[]

	// tag::expand[]
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets) //<1>
	// end::expand[]
}
