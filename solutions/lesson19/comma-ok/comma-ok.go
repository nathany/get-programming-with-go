package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temperature["Moon"] = 0

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vºC.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}
