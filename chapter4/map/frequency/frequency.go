package main

import "fmt"

func main() {
	// tag::main[]
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, t := range temperatures { //<1>
		frequency[t]++
	}

	for t, num := range frequency { //<2>
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
	// end::main[]
}
