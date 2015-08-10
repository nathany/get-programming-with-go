package main

import "fmt"

func main() {
	const distance = 24000000000000000000 //#A
	const lightSpeed = 299792
	const secondsPerDay = 86400

	days := distance / lightSpeed / secondsPerDay //#B

	fmt.Println("Andromeda Galaxy is", days, "light days away.") //#C

	fmt.Println("Andromeda Galaxy is", float64(distance), "km away.")
}

// #A Exceeds the range of a 64-bit integer
// #B The days variable is a regular int
// #C Andromeda Galaxy is 926568346 light days away.
