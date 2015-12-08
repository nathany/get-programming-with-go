package main

import "fmt"

func main() {
	// tag::main[]
	var board [8][8]string // <1>

	board[0][0] = "♜"
	board[0][7] = "♜" // <2>

	for column := range board[1] {
		board[1][column] = "♟"
	}

	fmt.Print(board)
	// end::main[]
}
