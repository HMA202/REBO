package main

import "fmt"

func main() {
	fmt.Println(CaesarCipher("abc XYZ", 2))
	fmt.Println(CaesarCipher("cde ZAB", -2))
}
