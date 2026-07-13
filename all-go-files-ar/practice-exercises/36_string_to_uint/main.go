package main

import "fmt"

func main() {
	fmt.Println(StringToUint("123"))
	fmt.Println(StringToUint("0"))
	fmt.Println(StringToUint("99"))
	fmt.Println(StringToUint("-5"))
}
