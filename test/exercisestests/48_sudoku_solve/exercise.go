package main

func sudokuCanPlace(board *[9][9]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == num {
				return false
			}
		}
	}

	return true
}

func sudokuFindEmptyCell(board *[9][9]int) (int, int, bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return row, col, true
			}
		}
	}

	return -1, -1, false
}

func SudokuSolve(board *[9][9]int) bool {
	row, col, found := sudokuFindEmptyCell(board)

	if !found {
		return true
	}

	for num := 1; num <= 9; num++ {
		if sudokuCanPlace(board, row, col, num) {
			board[row][col] = num

			if SudokuSolve(board) {
				return true
			}

			board[row][col] = 0
		}
	}

	return false
}
