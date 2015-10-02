package main

import "fmt"

func main() {
	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400

	days := distance / lightSpeed / secondsPerDay

	fmt.Println("Andromeda Galaxy is", days, "light days away.")

	fmt.Println("Andromeda Galaxy is", float64(distance), "km away.")
}
