package main

func main() {
	type celsius float64

	const degrees = 20
	var temperature celsius = degrees

	temperature += 10
}
