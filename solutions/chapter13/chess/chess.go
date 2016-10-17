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
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'

	// pawns
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}

	// white pieces
	board[7][0] = 'R'
	board[7][1] = 'N'
	board[7][2] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'
	board[7][5] = 'B'
	board[7][6] = 'N'
	board[7][7] = 'R'

	display(board)
}
