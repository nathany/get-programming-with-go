package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vºC.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	fmt.Println(temperature["Moon"])

	if temp, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vºC.\n", temp)
	} else {
		fmt.Println("Where is the moon?")
	}

	_, ok := temperature["Earth"]
	fmt.Println(ok)

	delete(temperature, "Earth")
}
