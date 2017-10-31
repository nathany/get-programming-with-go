package main

import (
	"fmt"
	"time"
)

func main() {
	const layout = "Mon, Jan 2, 2006"

	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)

	fmt.Println(day.Format(layout))
	fmt.Println(tomorrow.Format(layout))
}
