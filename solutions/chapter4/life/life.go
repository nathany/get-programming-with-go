package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Width  = 80
	Height = 15
)

// Universe is a two-dimensional field of cells.
type Universe [][]bool

// NewUniverse returns an empty universe.
func NewUniverse() Universe {
	u := make(Universe, Height)
	for i := range u {
		u[i] = make([]bool, Width)
	}
	return u
}

// Seed random live cells into the universe.
func (u Universe) Seed() {
	for i := 0; i < (Width * Height / 4); i++ {
		u.Set(rand.Intn(Width), rand.Intn(Height), true)
	}
}

// Set the state of the specified cell.
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the coordinates are outside of the universe, they wrap around.
func (u Universe) Alive(x, y int) bool {
	x = (x + Width) % Width
	y = (y + Height) % Height
	return u[y][x]
}

// Neighbors counts the adjacent cells that are alive.
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

// Next returns the state of the specified cell at the next step.
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

// String returns the universe as a string.
func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (Width+1)*Height)

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}

	return string(buf)
}

// Show clears the screen and displays the universe.
func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}

// Step updates the state of the next universe (b) from
// the current universe (a).
func Step(a, b Universe) {
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a // Swap universes
	}
}
