package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' ||
		r == 'e' ||
		r == 'i' ||
		r == 'o' ||
		r == 'u' ||
		r == 'A' ||
		r == 'E' ||
		r == 'I' ||
		r == 'O' ||
		r == 'U'
}

func main() {
	vowels := []rune{}

	for _, argument := range os.Args[1:] {
		for _, r := range argument {
			if isVowel(r) {
				vowels = append(vowels, r)
			}
		}
	}

	vowelIndex := len(vowels) - 1

	for argumentIndex, argument := range os.Args[1:] {
		for _, r := range argument {
			if isVowel(r) {
				z01.PrintRune(vowels[vowelIndex])
				vowelIndex--
			} else {
				z01.PrintRune(r)
			}
		}

		if argumentIndex < len(os.Args[1:])-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
