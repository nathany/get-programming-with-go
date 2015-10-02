package main

import "fmt"

func main() {
	// tag::main[]
	const distance = 24000000000000000000 //<1>
	const lightSpeed = 299792
	const secondsPerDay = 86400

	days := distance / lightSpeed / secondsPerDay //<2>

	fmt.Println("Andromeda Galaxy is", days, "light days away.") //<3>

	fmt.Println("Andromeda Galaxy is", float64(distance), "km away.")
	// end::main[]
}
