package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var queens [8]int

	placeQueen(&queens, 0)
}

func placeQueen(queens *[8]int, column int) {
	if column == len(queens) {
		printQueens(queens)
		return
	}

	row := 1
	for row <= len(queens) {
		if isSafePosition(queens, column, row) {
			queens[column] = row
			placeQueen(queens, column+1)
		}

		row++
	}
}

func isSafePosition(queens *[8]int, column int, row int) bool {
	previousColumn := 0

	for previousColumn < column {
		if queens[previousColumn] == row {
			return false
		}

		rowDifference := queens[previousColumn] - row
		if rowDifference < 0 {
			rowDifference = -rowDifference
		}

		columnDifference := column - previousColumn

		if rowDifference == columnDifference {
			return false
		}

		previousColumn++
	}

	return true
}

func printQueens(queens *[8]int) {
	digits := [8]rune{'1', '2', '3', '4', '5', '6', '7', '8'}
	index := 0

	for index < len(queens) {
		z01.PrintRune(digits[queens[index]-1])
		index++
	}

	z01.PrintRune('\n')
}
