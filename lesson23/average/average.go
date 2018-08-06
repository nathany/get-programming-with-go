package main

import "fmt"

type report struct {
	sol         int
	location    location
	temperature temperature
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}
func (r report) average() celsius {
	return r.temperature.average()
}

func main() {
	t := temperature{high: -1.0, low: -78.0}
	fmt.Printf("average %vº C\n", t.average())
	report := report{sol: 15, temperature: t}
	fmt.Printf("average %vº C\n", report.temperature.average())
}
