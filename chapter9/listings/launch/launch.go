package main

import "fmt"

func main() {
	launch := false

	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)

	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)
}
