package main

func main() {
	// tag::main[]
	type Celsius float64

	const degrees = 20
	var temperature Celsius = degrees

	temperature += 10
	// end::main[]
}
