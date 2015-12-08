package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	// tag::append[]
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"} //<1>
	dwarfs2 := append(dwarfs1, "Orcus")                                 //<2>
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")            //<3>
	// end::append[]

	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)

	// tag::mutate[]
	dwarfs3[1] = "Pluto!" //<1>
	// end::mutate[]

	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
}
