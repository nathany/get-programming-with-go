package main

import "fmt"

func main() {
	dwarfs := make([]string, 0, 10)

	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")

	fmt.Println(dwarfs)
}
