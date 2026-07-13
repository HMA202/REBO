package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, argument := range os.Args[1:] {
		for _, r := range argument {
			z01.PrintRune(r)
		}

		z01.PrintRune('\n')
	}
}
