package main

import (
	"fmt"
	"sort"
)

func main() {
	// tag::set[]
	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool) //<1>
	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("set member") //<2>
	}

	fmt.Println(set) //<3>
	// end::set[]

	// tag::slice[]
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)

	fmt.Println(unique) //<1>
	// end::slice[]
}
