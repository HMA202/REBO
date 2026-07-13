package main

import "fmt"

func main() {
	fmt.Println(ParseBool("true"))
	fmt.Println(ParseBool("false"))
	fmt.Println(ParseBool("yes"))
}
