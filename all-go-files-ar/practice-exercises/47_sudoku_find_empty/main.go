package main

import "fmt"

func main() {
	board := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
	}

	row, col, found := SudokuFindEmpty(board)
	fmt.Println(row)
	fmt.Println(col)
	fmt.Println(found)
}
