package main

import "fmt"

func main() {
	yesNo := "1"

	var launch bool

	switch yesNo {
	case "true", "yes", "1":
		launch = true
	case "false", "no", "0":
		launch = false
	default:
		fmt.Println(yesNo, "is not valid")
	}

	fmt.Println("Ready for launch:", launch)
}
