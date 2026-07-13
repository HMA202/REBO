package main

import "fmt"

func PrintSquare(n int) {
	if n <= 0 {
		return
	}
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
