package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// tag::main[]
	type point struct {
		Lat  float64 `json:"latitude"`  //<1>
		Long float64 `json:"longitude"` //<1>
	}

	curiosity := point{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes)) //<2>
	// end::main[]
}

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
