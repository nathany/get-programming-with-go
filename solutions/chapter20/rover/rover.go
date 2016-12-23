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

type rover string

func (r rover) talk() string {
	return string(r)
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("toot ", int(l))
}

func main() {
	r := rover("whir whir")
	shout(r)

	shout(laser(2))
}
