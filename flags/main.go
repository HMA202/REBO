package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	fmt.Print("--insert\n")
	fmt.Print("  -i\n")
	fmt.Print("\t This flag inserts the string into the string passed as argument.\n")
	fmt.Print("--order\n")
	fmt.Print("  -o\n")
	fmt.Print("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}

	return s[:len(prefix)] == prefix
}

func sortASCII(s string) string {
	chars := []rune(s)

	for i := 0; i < len(chars)-1; i++ {
		for j := 0; j < len(chars)-1-i; j++ {
			if chars[j] > chars[j+1] {
				chars[j], chars[j+1] = chars[j+1], chars[j]
			}
		}
	}

	return string(chars)
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 1 {
		printHelp()
		return
	}

	insert := ""
	order := false
	text := ""

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		}

		if arg == "--order" || arg == "-o" {
			order = true
			continue
		}

		if hasPrefix(arg, "--insert=") {
			insert = arg[len("--insert="):]
			continue
		}

		if hasPrefix(arg, "-i=") {
			insert = arg[len("-i="):]
			continue
		}

		if arg == "--insert" || arg == "-i" {
			if i+1 < len(os.Args) {
				i++
				insert = os.Args[i]
			}
			continue
		}

		text = arg
	}

	result := text + insert

	if order {
		result = sortASCII(result)
	}

	printString(result)
}
