package main

import (
	"fmt"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go removeDuplicates(c0, c1)
	printGopher(c1)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"a", "b", "b", "c", "d", "d", "d", "e"} {
		downstream <- v
	}
	close(downstream)
}

func removeDuplicates(upstream, downstream chan string) {
	prev := ""
	for v := range upstream {
		if v != prev {
			downstream <- v
			prev = v
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
