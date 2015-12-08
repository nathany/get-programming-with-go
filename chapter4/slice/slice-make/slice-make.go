package main

import "fmt"

func main() {
	// tag::main[]
	dwarfs := make([]string, 0, 10)

	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	// end::main[]

	fmt.Println(dwarfs)
}
