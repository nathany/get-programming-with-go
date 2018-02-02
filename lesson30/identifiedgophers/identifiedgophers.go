package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}

func sleepyGopher(whichGopher int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", whichGopher, " snore ...")
}
