package piscine

import "github.com/01-edu/z01"

func printRune(r rune) {
	z01.PrintRune(r)
}

func printLine(x int, left rune, middle rune, right rune) {
	for col := 1; col <= x; col++ {
		if col == 1 {
			printRune(left)
		} else if col == x {
			printRune(right)
		} else {
			printRune(middle)
		}
	}
	printRune('\n')
}

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		if row == 1 || row == y {
			printLine(x, 'o', '-', 'o')
		} else {
			printLine(x, '|', ' ', '|')
		}
	}
}

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		if row == 1 {
			printLine(x, '/', '*', '\\')
		} else if row == y {
			printLine(x, '\\', '*', '/')
		} else {
			printLine(x, '*', ' ', '*')
		}
	}
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		if row == 1 {
			printLine(x, 'A', 'B', 'A')
		} else if row == y {
			printLine(x, 'C', 'B', 'C')
		} else {
			printLine(x, 'B', ' ', 'B')
		}
	}
}

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		if row == 1 {
			printLine(x, 'A', 'B', 'C')
		} else if row == y {
			printLine(x, 'A', 'B', 'C')
		} else {
			printLine(x, 'B', ' ', 'B')
		}
	}
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		if row == 1 {
			printLine(x, 'A', 'B', 'C')
		} else if row == y {
			printLine(x, 'C', 'B', 'A')
		} else {
			printLine(x, 'B', ' ', 'B')
		}
	}
}
