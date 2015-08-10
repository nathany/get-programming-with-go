package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	// Alpha Centauri is 41.3 trillion kilometers away
	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.") //#A

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.") //#B
}

// #A Alpha Centauri is 41300000000000 km away.
// #B That is 1594 days of travel at light speed.
