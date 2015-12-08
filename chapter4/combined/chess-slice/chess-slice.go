package main

func main() {
	var board [][]string

	board = make([][]string, 8)
	for row := range board {
		board[row] = make([]string, 8)
	}

	board[0][0] = "♜"
	board[0][7] = "♜"
}
