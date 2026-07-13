package main

import "fmt"

func main() {
	board := [3][3]rune{
		{'X', 'X', 'X'},
		{'O', ' ', 'O'},
		{' ', ' ', ' '},
	}

	fmt.Println(TicTacToeWinner(board))
}
