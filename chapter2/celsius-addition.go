package main

func main() {
	type Celsius float64

	const degrees = 20
	var temperature Celsius = degrees

	temperature += 10
}
