package main

import "fmt"

func main() {
	// tag::main[]
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// end::main[]

	fmt.Println(dwarfSlice, dwarfs)
}
