package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vº C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)
	moon := temperature["Moon"]
	fmt.Println(moon)
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vº C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
	_, ok := temperature["Earth"]
	fmt.Println(ok)

	delete(temperature, "Earth")
}
