package main

import "sync"

// Visited tracks whether web pages have been visited.
// Its methods may be used concurrently with one another.
type Visited struct {
	// mu guards the visited map.
	mu      sync.Mutex
	visited map[string]int
}

// VisitLink tracks that the page with the given URL has
// been visited, and returns the updated link count.
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
}
