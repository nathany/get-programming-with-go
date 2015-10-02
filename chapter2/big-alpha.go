// Alpha Centauri is 41.3 trillion kilometers away
package main

import "fmt"

func main() {
	// tag::main[]
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.") //<1>

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.") //<2>
	// end::main[]
}
