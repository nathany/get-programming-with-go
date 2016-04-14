package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")

	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
	dwarfs3[1] = "Pluto!"

	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
}
