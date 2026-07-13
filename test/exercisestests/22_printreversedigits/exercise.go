package main

import "fmt"

func PrintReverseDigits() {
	for i := 9; i >= 0; i-- {
		fmt.Print(i)
	}
	fmt.Println()
}
