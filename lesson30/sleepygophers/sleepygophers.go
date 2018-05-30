package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	time.Sleep(4 * time.Second)
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}
