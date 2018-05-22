package main

import "fmt"

func main() {
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Println(dwarfSlice, dwarfs)
	fmt.Printf("array %T\n", dwarfArray)
	fmt.Printf("slice %T\n", dwarfs)
}
