package main

import "fmt"

func main() {
	for count := 10; count > 0; count-- { //#A
		fmt.Println("T minus", count)
	} //#B
	fmt.Println("Liftoff!")

	for count := 1; count <= 10; count++ { //#C
		fmt.Println("T plus", count)
	} //#D
}

// #A The count variable is declared and in scope
// #B Count is no longer in scope
// #C A new variable named count is declared
// #D The second count variable is no longer in scope
