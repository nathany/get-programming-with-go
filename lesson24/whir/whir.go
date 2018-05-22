package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("toot ", int(l))
}

type rover string

func (r rover) talk() string {
	return string(r)
}

func main() {
	r := rover("whir whir")
	shout(r)
}
