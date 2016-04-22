package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type point struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	type dataPoint struct {
		Sol      int   `json:"sol"`
		Location point `json:"location"`
	}

	curiosity := point{-4.5895, 137.4417}
	data := dataPoint{Sol: 0, Location: curiosity}

	bytes, err := json.Marshal(data)
	exitOnError(err)

	fmt.Println(string(bytes))
}

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
