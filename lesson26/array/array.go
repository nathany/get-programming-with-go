package main

import "fmt"

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
	// ...
}

func main() {
	var board [8][8]rune
	reset(&board)

	fmt.Printf("%c", board[0][0])
}
