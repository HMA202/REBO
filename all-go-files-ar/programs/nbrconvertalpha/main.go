package main

import (
	"os"

	"github.com/01-edu/z01"
)

func getPosition(s string) int {
	if len(s) == 0 {
		return 0
	}

	number := 0

	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}

		number = number*10 + int(r-'0')

		if number > 26 {
			return 0
		}
	}

	return number
}

func main() {
	upper := false
	start := 1

	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		upper = true
		start = 2
	}

	if start >= len(os.Args) {
		return
	}

	for i := start; i < len(os.Args); i++ {
		position := getPosition(os.Args[i])

		if position < 1 || position > 26 {
			z01.PrintRune(' ')
		} else if upper {
			z01.PrintRune('A' + rune(position-1))
		} else {
			z01.PrintRune('a' + rune(position-1))
		}
	}

	z01.PrintRune('\n')
}
