package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	digits := [10]int{}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	for n > 0 {
		digits[n%10]++
		n /= 10
	}

	for digit := 0; digit <= 9; digit++ {
		for digits[digit] > 0 {
			z01.PrintRune(rune('0' + digit))
			digits[digit]--
		}
	}
}
