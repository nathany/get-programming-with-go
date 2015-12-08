package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type point struct {
		Lat, Long float64 //<1>
	}

	curiosity := point{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes)) //<2>
}

// exitOnError prints any errors and exits.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
