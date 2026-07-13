package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	comb := make([]int, n)
	first := true

	var loop func(pos int, start int)

	loop = func(pos int, start int) {
		if pos == n {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

			for i := 0; i < n; i++ {
				z01.PrintRune('0' + rune(comb[i]))
			}

			first = false
			return
		}

		for i := start; i <= 10-n+pos; i++ {
			comb[pos] = i
			loop(pos+1, i+1)
		}
	}

	loop(0, 0)
	z01.PrintRune('\n')
}
