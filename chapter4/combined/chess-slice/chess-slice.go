package main

func main() {
	// tag::main[]
	var board [][]string //<1>

	board = make([][]string, 8) //<2>
	for row := range board {
		board[row] = make([]string, 8) //<3>
	}

	board[0][0] = "♜"
	board[0][7] = "♜"
	// end::main[]
}
