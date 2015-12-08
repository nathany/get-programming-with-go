package main

import "fmt"

func display(board [8][8]rune) {
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", column)
			}
		}
		fmt.Println()
	}
}

func main() {
	var board [8][8]rune

	// black pieces
	board[0][0] = '♜'
	board[0][1] = '♞'
	board[0][2] = '♝'
	board[0][3] = '♛'
	board[0][4] = '♚'
	board[0][5] = '♝'
	board[0][6] = '♞'
	board[0][7] = '♜'

	// pawns
	for column := range board[1] {
		board[1][column] = '♟'
		board[6][column] = '♙'
	}

	// white pieces
	board[7][0] = '♖'
	board[7][1] = '♘'
	board[7][2] = '♗'
	board[7][3] = '♕'
	board[7][4] = '♔'
	board[7][5] = '♗'
	board[7][6] = '♘'
	board[7][7] = '♖'

	display(board)
}
