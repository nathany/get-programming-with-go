package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

// Grid is a Sudoku grid
type Grid [rows][columns]int8

// Set a digit on a Sudoku grid
func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("out of bounds")
	}

	g[row][column] = digit
	return nil
}
func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}
func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		fmt.Printf("An error occurred: %v.\n", err)
		os.Exit(1)
	}
}
