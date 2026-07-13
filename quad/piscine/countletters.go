package piscine

import "github.com/01-edu/z01"

func printNumber(n int) {
	if n >= 10 {
		printNumber(n / 10)
	}
	z01.PrintRune('0' + rune(n%10))
}

func CountLetters(s string) {
	i := 0

	for i < len(s) {
		if s[i] == ' ' {
			z01.PrintRune(' ')
			i++
			continue
		}

		count := 1

		for i+count < len(s) && s[i+count] == s[i] {
			count++
		}

		printNumber(count)
		z01.PrintRune(rune(s[i]))

		i += count
	}

	z01.PrintRune('\n')
}
