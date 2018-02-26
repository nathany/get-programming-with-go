package main

import (
	"fmt"
	"time"
)

func main() {
	go worker()
	time.Sleep(5 * time.Second)
}
func worker() {
	n := 0
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			n++
			fmt.Println(n)
			next = time.After(time.Second)
		}
	}
}
