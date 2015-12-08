package main

import "fmt"

type StringSlice []string

func (slice StringSlice) dump(label string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := StringSlice{"Ceres", "Pluto", "Haumea"}
	dwarfs.dump("dwarfs") //<1>
}
