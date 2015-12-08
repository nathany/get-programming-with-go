package main

import "fmt"

func main() {
	// tag::main[]
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs) // <1>
	// end::main[]

	// tag::append[]
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs) // <1>
	// end::append[]
}
