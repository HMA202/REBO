package main

import "fmt"

func PrintOddDigits() {
	for i := 0; i <= 9; i++ {
		if i%2 != 0 {
			fmt.Print(i)
		}
	}
	fmt.Println()
}
