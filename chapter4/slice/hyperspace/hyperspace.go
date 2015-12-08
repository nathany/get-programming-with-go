package main

import (
	"fmt"
	"strings"
)

// hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) { //<1>
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" Venus   ", "Earth  ", " Mars"} //<2>
	hyperspace(planets)

	fmt.Println(strings.Join(planets, "")) //<3>
}
