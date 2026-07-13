package main

func SudokuFindEmpty(board [9][9]int) (int, int, bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return row, col, true
			}
		}
	}

	return -1, -1, false
}
