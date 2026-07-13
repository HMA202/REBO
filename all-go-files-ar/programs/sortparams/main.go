package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-1-i; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	for _, argument := range args {
		for _, r := range argument {
			z01.PrintRune(r)
		}

		z01.PrintRune('\n')
	}
}
