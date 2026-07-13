package main

func TicTacToeWinner(board [3][3]rune) string {
	lines := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range lines {
		a := board[line[0][0]][line[0][1]]
		b := board[line[1][0]][line[1][1]]
		c := board[line[2][0]][line[2][1]]

		if a != ' ' && a == b && b == c {
			return string(a)
		}
	}

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == ' ' {
				return "Pending"
			}
		}
	}

	return "Draw"
}
