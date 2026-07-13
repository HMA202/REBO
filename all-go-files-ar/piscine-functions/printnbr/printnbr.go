package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')

		if n <= -10 {
			PrintNbr(-(n / 10))
		}

		digit := -(n % 10)
		r := '0'
		for i := 0; i < digit; i++ {
			r++
		}
		z01.PrintRune(r)
		return
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}

	digit := n % 10
	r := '0'
	for i := 0; i < digit; i++ {
		r++
	}
	z01.PrintRune(r)
}
