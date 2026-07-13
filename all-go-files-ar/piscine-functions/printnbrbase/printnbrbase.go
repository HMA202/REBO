package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	baseRunes := []rune(base)

	if !validPrintNbrBase(baseRunes) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == 0 {
		z01.PrintRune(baseRunes[0])
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
	} else {
		nbr = -nbr
	}

	printNegativeNbrBase(nbr, baseRunes)
}

func validPrintNbrBase(base []rune) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}

		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func printNegativeNbrBase(nbr int, base []rune) {
	baseLength := len(base)

	if nbr <= -baseLength {
		printNegativeNbrBase(nbr/baseLength, base)
	}

	index := -(nbr % baseLength)
	z01.PrintRune(base[index])
}
