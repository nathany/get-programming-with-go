package main

import "sync"

var mu sync.Mutex

func main() {
	mu.Lock()
	defer mu.Unlock()
	// The lock is held until we return from the function.
}
