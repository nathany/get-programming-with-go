package main

import (
	"image"
	"sync"
)

func main() {
}

// MarsGrid represents a grid of some of the surface
// of Mars. It may be used concurrently by different
// goroutines.
type MarsGrid struct {
	bounds image.Rectangle
	mu     sync.Mutex
	cells  [][]cell
}

// cell holds information about a given cell in
// the grid.
type cell struct {
	occupier *Occupier
}

// NewMarsGrid returns a new MarsGrid of the
// given size.
func NewMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{
			Max: size,
		},
		cells: make([][]cell, size.Y),
	}
	for i := range grid.cells {
		grid.cells[i] = make([]cell, size.X)
	}
	return grid
}

// Size returns the size of the grid.
func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
}

// Occupy occupies a cell at the given point in the grid. It
// returns nil if the point is already or the point is outside
// the grid. Otherwise it returns a value that can be used
// to move to different places on the grid.
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	cell := g.cell(p)
	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		pos:  p,
	}
	return cell.occupier
}

// cell returns a pointer to the cell at the
// given position, or nil if it's outside
// the grid.
func (g *MarsGrid) cell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.Y][p.X]
}

// Occupier represents an occupied cell in the grid.
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

// MoveTo moves the occupier to a different cell in the grid.
// It reports whether the move was successful
// It might fail because it was trying to move outside
// the grid or because the cell it's trying to move into
// is occupied. If it fails, the occupier remains in the same place.
func (o *Occupier) MoveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	newCell := o.grid.cell(p)
	if newCell == nil || newCell.occupier != nil {
		return false
	}
	o.grid.cell(o.pos).occupier = nil
	newCell.occupier = o
	o.pos = p
	return true
}

// Pos returns the current grid position of the occupier.
func (o *Occupier) Pos() image.Point {
	return o.pos
}
