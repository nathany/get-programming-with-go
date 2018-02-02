package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 10; i++ {
		whichGopher := <-c
		fmt.Println("gopher ", whichGopher, " has finished sleeping")
	}
}

func sleepyGopher(whichGopher int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", whichGopher, " snore ...")
	c <- whichGopher
}
