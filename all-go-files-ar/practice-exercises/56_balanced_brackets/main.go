package main

import "fmt"

func main() {
	fmt.Println(BalancedBrackets("([]{})"))
	fmt.Println(BalancedBrackets("([)]"))
	fmt.Println(BalancedBrackets("((hello))"))
}
