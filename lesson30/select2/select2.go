package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go sleepyGopher(i, c)
	}
	timeout := time.After(2 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case whichGopher := <-c:
			fmt.Println("gopher ", whichGopher, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}
func sleepyGopher(whichGopher int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- whichGopher
}
