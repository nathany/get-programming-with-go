package main

import "fmt"

func main() {
	var board [8][8]string

	board[0][0] = "♜"
	board[0][7] = "♜"

	for column := range board[1] {
		board[1][column] = "♟"
	}

	fmt.Print(board)
}
