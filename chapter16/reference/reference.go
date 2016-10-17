package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
}
