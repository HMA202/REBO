package main

func SudokuIsValid(board [9][9]int) bool {
	for row := 0; row < 9; row++ {
		seen := [10]bool{}

		for col := 0; col < 9; col++ {
			n := board[row][col]

			if n < 0 || n > 9 {
				return false
			}

			if n != 0 {
				if seen[n] {
					return false
				}
				seen[n] = true
			}
		}
	}

	for col := 0; col < 9; col++ {
		seen := [10]bool{}

		for row := 0; row < 9; row++ {
			n := board[row][col]

			if n != 0 {
				if seen[n] {
					return false
				}
				seen[n] = true
			}
		}
	}

	for boxRow := 0; boxRow < 9; boxRow += 3 {
		for boxCol := 0; boxCol < 9; boxCol += 3 {
			seen := [10]bool{}

			for row := boxRow; row < boxRow+3; row++ {
				for col := boxCol; col < boxCol+3; col++ {
					n := board[row][col]

					if n != 0 {
						if seen[n] {
							return false
						}
						seen[n] = true
					}
				}
			}
		}
	}

	return true
}
