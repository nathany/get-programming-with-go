package main

import (
	"fmt"
	"strings"
)

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	type starship struct {
		laser
	}

	s := starship{laser(3)}

	fmt.Println(s.talk())
	shout(s)
}
