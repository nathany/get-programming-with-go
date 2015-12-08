package main

import "fmt"

func main() {
	// tag::main[]
	temperature := map[string]int{
		"Earth": 15, // <1>
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vºC.\n", temp) //<2>

	temperature["Earth"] = 16 //<3>
	temperature["Venus"] = 464

	fmt.Println(temperature) //<4>
	// end::main[]

	// tag::moon[]
	fmt.Println(temperature["Moon"]) //<1>

	if temp, ok := temperature["Moon"]; ok { // <2>
		fmt.Printf("On average the moon is %vºC.\n", temp)
	} else {
		fmt.Println("Where is the moon?") //<3>
	}
	// end::moon[]

	// tag::delete[]
	_, ok := temperature["Earth"]
	fmt.Println(ok) //<1>

	delete(temperature, "Earth") //<2>
	// end::delete[]
}
