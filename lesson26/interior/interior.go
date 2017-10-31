package main

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

func main() {
	type character struct {
		name  string
		stats stats
	}

	player := character{name: "Matthias"}
	levelUp(&player.stats)

	fmt.Printf("%+v\n", player.stats)
}
